open Lwt

module BlogPost = struct
  type t = {
      slug : string;
      title : string;
      content : string;
      date : string;
  }
end

module Q = struct
  open Caqti_request.Infix
  open Caqti_type.Std

  let blog_post =
    let open BlogPost in
    let encode {slug; title; content; date} = Ok (slug, title, content, date) in
    let decode (slug, title, content, date) = Ok {slug; title; content; date} in
    let b_rep = Caqti_type.(tup4 string string string string) in
    custom ~encode ~decode b_rep

  let create_blog_post_table =
    unit ->. unit @@
      {eos|
       CREATE TABLE IF NOT EXISTS blog_post (
       slug text PRIMARY KEY,
       title text NOT NULL,
       content text NOT NULL,
       date text NOT NULL,
       active TINYINT DEFAULT 1,
       hidden_date DEFAULT (unixepoch('now'))
       |eos}
  
  let create_blog_post =
    tup4 string string string string ->. unit @@
      "INSERT INTO blog_post (slug, title, content, date) VALUES (?, ?, ?, ?)"

  let get_all_blog_posts =
    unit ->* blog_post @@
      "SELECT slug, title, content, date FROM blog_post WHERE active = 1 ORDER BY hidden_date DESC"

  let get_all_blog_posts_for_display =
    unit ->* tup2 string string @@
      "SELECT slug, title FROM blog_post WHERE active = 1"

  let get_blog_post_by_slug =
    string ->? blog_post @@
      "SELECT slug, title, content, date FROM blog_post WHERE slug = ? AND active = 1"

  let update_blog_post_content =
    tup2 string string ->. unit @@
      "UPDATE blog_post SET content = ? WHERE slug = ?"

  let update_blog_post_title =
    tup2 string string ->. unit @@
      "UPDATE blog_post SET title = ? WHERE slug = ?"

  let delete_blog_post =
    string ->. unit @@
      "UPDATE blog_post SET active = 0 WHERE slug = ?"
end

module Db = (val Caqti_lwt.connect (Uri.of_string "sqlite3:database.db") >>= Caqti_lwt.or_fail |> Lwt_main.run)

(* Wrappers for the generic functions defined in Q *)
module Database = struct
  let create_blog_post_table () =
    Db.exec Q.create_blog_post_table ()

  let create_blog_post slug title content date =
    Db.exec Q.create_blog_post (slug, title, content, date)

  let update_blog_post_content slug content =
    Db.exec Q.update_blog_post_content (content, slug)

  let update_blog_post_title slug title =
    Db.exec Q.update_blog_post_title (title, slug)

  let get_blog_post_by_slug slug =
    Db.find_opt Q.get_blog_post_by_slug slug

  let get_all_blog_posts () =
    Db.collect_list Q.get_all_blog_posts ()

  let iter_blog_posts f =
    Db.iter_s Q.get_all_blog_posts f ()

  let (>>=?) monad func =
    monad >>= (function | Ok x -> func x | Error err -> Lwt.return (Error err))

  let report_error = function
    | Ok () -> Lwt.return_unit
    | Error err ->
       Lwt_io.eprintl (Caqti_error.show err)
end
