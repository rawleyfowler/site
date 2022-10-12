module Blog_post :
  sig
    type t = {
      slug : string;
      title : string;
      content : string;
      date : string;
    }
  end

module Db : Caqti_lwt.CONNECTION

val create_blog_post_table :
  unit -> (unit, [> Caqti_error.call_or_retrieve ]) Stdlib.result Lwt.t

val create_blog_post :
  string ->
  string ->
  string ->
  string -> (unit, [> Caqti_error.call_or_retrieve ]) Stdlib.result Lwt.t

val update_blog_post_content :
  string ->
  string -> (unit, [> Caqti_error.call_or_retrieve ]) Stdlib.result Lwt.t

val update_blog_post_title :
  string ->
  string -> (unit, [> Caqti_error.call_or_retrieve ]) Stdlib.result Lwt.t

val get_blog_post_by_slug :
  string ->
  (Blog_post.t option, [> Caqti_error.call_or_retrieve ]) Stdlib.result Lwt.t

val get_all_blog_posts :
  unit ->
  (Blog_post.t list, [> Caqti_error.call_or_retrieve ]) Stdlib.result Lwt.t

val iter_blog_posts :
  (Blog_post.t ->
    (unit, [> Caqti_error.call_or_retrieve ] as 'a) Stdlib.result Lwt.t) ->
  (unit, 'a) Stdlib.result Lwt.t

val ( >>=? ) :
  ('a, 'b) Stdlib.result Lwt.t ->
  ('a -> ('c, 'b) Stdlib.result Lwt.t) -> ('c, 'b) Stdlib.result Lwt.t

val report_error : (unit, [< Caqti_error.t ]) Stdlib.result -> unit Lwt.t
