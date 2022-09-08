open Lwt.Infix
open Database
open Render

let db_created = ref true

let _ = Database.create_blog_post_table () >>= function
  | Ok () -> Lwt_io.print "Database initialized successfully.\n"
  | Error _ -> Lwt.return (db_created := false)

let () = Unix.sleepf 2.0
let () = if not !db_created then failwith "Could not initialize database"

let () =
  Dream.run
  @@ Dream.logger
  @@ Dream.router [
         Dream.get "/static/**" @@ Dream.static "static";
         Dream.get "/" @@ Dream.from_filesystem "html" "index.html";
         Dream.get "/resume" @@ Dream.from_filesystem "html" "resume.html";
         Dream.get "/philosophy" @@ Dream.from_filesystem "html" "philosophy.html";
         Dream.get "/web-ring" @@ Dream.from_filesystem "html" "web-ring.html";
         Dream.get "/blog" Render.render_blog_index;
         Dream.scope "/blog" [Dream.origin_referrer_check] [
             Dream.get "/:post" Render.render_blog_post
           ];
       ]
