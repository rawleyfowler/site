open Lwt.Infix
open Database
open Render

exception DatabaseCreationFailure

let _ =
  Database.create_blog_post_table () >>= function
  | Ok () -> Lwt_io.print "Database was initialized successfully\n"
  | Error _ -> raise DatabaseCreationFailure

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
