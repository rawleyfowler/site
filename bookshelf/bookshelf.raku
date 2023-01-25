use v6.d;

use Humming-Bird::Core;
use Template::Mustache;

my constant $BOOKSHELF-LOCATION = 'bookshelf.txt';

my $tmpl = Template::Mustache.new: :from<.>;

get('/bookshelf', -> $request, $response {
  my @books = $BOOKSHELF-LOCATION.IO.lines>>.split('^').map(-> ($name, $author, $href, $rating) { %(:$name, :$author, :$href, :$rating) });
  $response.html($tmpl.render('bookshelf', { :@books }));
});

listen(8888);
