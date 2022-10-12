module Blog_post = Database.Blog_post
module type SimpleRender = sig val render : unit -> string end
val header_template : string
val footer_template : string
val replace_sequence : string -> string -> string -> string
val html_unescape : string -> string
val not_found_template : string
val error_template : string
val render_page : string -> string
val generate_link : Blog_post.t -> string
val generate_rss_item : Blog_post.t -> string
val handle_error : [< Caqti_error.t ] -> Dream.response Lwt.t
val handle_not_found : unit -> Dream.response Lwt.t
val render_blog_index : unit -> Dream.response Lwt.t
val render_rss_feed : Dream.request -> Dream.response Lwt.t
val render_simple : (module SimpleRender) -> Dream.response Lwt.t
val render_blog_post : slug:string -> Dream.response Lwt.t
val render_index : unit -> Dream.response Lwt.t
