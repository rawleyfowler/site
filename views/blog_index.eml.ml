let render ~posts =
  let open Database in
  <div>
  <h3>Blog</h3>
  I have an <a href="/blog/rss.xml">RSS feed</a> too.
  <br>
  </div>
% posts |> List.iter begin fun (post : Blog_post.t) ->
  <div class="link-wrapper">
    <a href="/blog/<%s post.slug %>">
      <%s post.title %>
    </a>
    <i>
      <%s post.date %>
    </i>
  </div>
% end;