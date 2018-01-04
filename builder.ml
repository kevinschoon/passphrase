open Core

let _ = Random.self_init()

module Password : sig 
  type t
  val of_strings : string list -> t
  val to_string : t -> string
end = struct 
  type t = string list
  let of_strings words = words
  let to_string t = String.concat t
end

let add_word entries line =
  let word = List.nth_exn (String.split ~on:' ' line) 1 in
  word :: entries

let entries source = 
  List.fold ~init:[] ~f: add_word (String.split ~on:'\n' source)

let get_word entries = 
  let count = List.length entries in
  let role = Random.int count in
  let entry = List.nth_exn entries role in
    entry

let build count =
  let e = entries Words.long in
    Password.of_strings (List.init count ~f:(fun i -> get_word e))
