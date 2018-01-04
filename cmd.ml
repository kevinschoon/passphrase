open Core

let command = 
  Command.basic' 
  ~summary: "Generate random word-based passphrases"
  ~readme: (fun () -> "Passphrase generates a diceware-style random passphrase based on wordlists published by EFF.")
  (let open Command.Let_syntax in
  [%map_open
    (* flags *)
    let words = flag "-words" (optional int) ~doc: "number of words to generate"
    in
    fun () ->
      let password = Builder.build (match words with | None -> 4 | Some words -> words) in
      printf "%s\n" (Builder.Password.to_string password)
  ])

let () =
  Command.run ~version: "0.1" ~build_info:"RWO" command
