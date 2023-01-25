#!/bin/bash
# This is free-software, and a part of github.com/rawleyfowler/rawleydotxyz
# NOTE: This will most likely be replaced by a Perl or Raku script at some point.
function usage() {
	echo "Usage: insert-posts"
	echo "Tool for inserting posts into a SQLite3 database."
	echo 'Posts should be located in $BLOG_DIR/posts, and the database should be in $BLOG_DIR'
	exit 1
}

[ ! -x "$(command -v sqlite3)" ] && echo "You need SQLite3 to run this" && exit 1
if [ -z "$BLOG_DIR" ]; then
    echo "Looking for posts in ./posts because BLOG_DIR is not set."
    BLOG_DIR="."
fi

# params -> slug, title, content, date
function update_post() {
    local date=${1:?Date is required}
    local slug=${2:?Slug is required}
    local title=${3:?Title is required}
    local content=${4:?Content is required}

    echo "Trying to update $slug"

    sqlite3 \
	"$BLOG_DIR/database.db" \
	"UPDATE blog_post SET title = '$title', content = '$content' WHERE slug = '$slug'" 2> /dev/null

    local status="$?"

    [[ "$status" = "0" ]] && echo "Successfully updated already existing post: $slug" && return
    >&2 echo "Failed to update, or insert post: $slug."
}

# params -> slug, title, content, date
function insert_post() {
    [ ! $# -eq 4 ] && echo "something went wrong" && exit 1

    local date=${1:?Date is required}
    local slug=${2:?Slug is required}
    local title=${3:?Title is required}
    local content=${4:?Content is required}
    
    sqlite3 "$BLOG_DIR/database.db" "INSERT INTO blog_post (slug, title, content, date) VALUES ('$slug', '$title', '$content', '$date');" 2> /dev/null

    local status="$?"

    [ "$status" = "0" ] && echo "Inserted $slug successfully." && return
    [ "$status" = "8" ] && echo "The database is read restricted, did you run the program as the correct user?" && exit 1
    [ "$status" = "19" ] && echo "Post $slug already exists." && update_post "$date" "$slug" "$title" "$content"
}

# # # # # # # # # # # # # # # # # # # # # # # #
# Blog post files are structured as follows:  #
# FIRST LINE -> date in desired format        #
# SECOND LINE -> URI friendly slug            #
# THIRD LINE -> Title of post                 #
# REMAINING LINES -> Content of post          #
# # # # # # # # # # # # # # # # # # # # # # # #

for file in $(ls "$BLOG_DIR/posts"); do
    echo "Trying to insert: $file"

    target_file="$BLOG_DIR/posts/$file"
    date=$(sed -n '1p' $target_file | tr -d '\n')
    slug=$(sed -n '2p' $target_file | tr -d '\n')
    title=$(sed -n '3p' $target_file | tr -d '\n')
    content=$(sed -n '4,$p' $target_file | tr -d '\n')

    content=${content//" & "/&amp;}
    content=${content//"'"/&lsquo;}

    insert_post "$date" "$slug" "$title" "$content"
done
