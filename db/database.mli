module BlogPost :
  sig
    type t = {
      slug : string;
      title : string;
      content : string;
      date : string;
    }
  end

module Q :
  sig
    val blog_post : BlogPost.t Caqti_type.t
    
    val create_blog_post_table : (unit, unit, [ `Zero ]) Caqti_request.t
    
    val create_blog_post :
      (string * string * string * string, unit, [ `Zero ]) Caqti_request.t
    
    val get_all_blog_posts :
      (unit, BlogPost.t, [ `Many | `One | `Zero ]) Caqti_request.t
    
    val get_all_blog_posts_for_display :
      (unit, string * string, [ `Many | `One | `Zero ]) Caqti_request.t
    
    val get_blog_post_by_slug :
      (string, BlogPost.t, [ `One | `Zero ]) Caqti_request.t
    
    val update_blog_post_content :
      (string * string, unit, [ `Zero ]) Caqti_request.t
    
    val update_blog_post_title :
      (string * string, unit, [ `Zero ]) Caqti_request.t
    
    val delete_blog_post : (string, unit, [ `Zero ]) Caqti_request.t
  end

module Db : Caqti_lwt.CONNECTION

module Database :
  sig
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
      (BlogPost.t option, [> Caqti_error.call_or_retrieve ]) Stdlib.result Lwt.t
    
    val get_all_blog_posts :
      unit ->
      (BlogPost.t list, [> Caqti_error.call_or_retrieve ]) Stdlib.result Lwt.t
    
    val iter_blog_posts :
      (BlogPost.t ->
       (unit, [> Caqti_error.call_or_retrieve ] as 'a) Stdlib.result Lwt.t) ->
      (unit, 'a) Stdlib.result Lwt.t
    
    val ( >>=? ) :
      ('a, 'b) Stdlib.result Lwt.t ->
      ('a -> ('c, 'b) Stdlib.result Lwt.t) -> ('c, 'b) Stdlib.result Lwt.t
    
    val report_error : (unit, [< Caqti_error.t ]) Stdlib.result -> unit Lwt.t
  end
