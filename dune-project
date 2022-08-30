(lang dune 3.4)

(name rawleydotxyz)
(generate_opam_files true)

(source
	(github rawleyfowler/rawleydotxyz))
(authors "Rawley Fowler")
(maintainers "Rawley Fowler")
(license "AGPL3.0")

(package
	(name rawleydotxyz)
	(synopsis "Rawley Fowler's website + blog")
	(description "Rawley Fowler's personal website and blog, www.rawley.xyz")
	(allow_empty)
	(depends
	  ocaml
	  dune
	  lwt
	  dream
	  caqti
	  caqti-driver-sqlite))