open Database
  
module Render :
  sig
    val not_found_template : string
    val error_template : string
    val header_template : string
    val footer_template : string
    val render_page : string -> string
    val generate_link : BlogPost.t -> string
    val handle_error : [< Caqti_error.t ] -> Dream.response Lwt.t
    val handle_not_found : unit -> Dream.response Lwt.t
    val render_blog_post : Dream.request -> Dream.response Lwt.t
    val render_blog_index : Dream.request -> Dream.response Lwt.t
    val render_rss_feed : Dream.request -> Dream.response Lwt.t
  end
