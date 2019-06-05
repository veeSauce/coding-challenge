# key-value store
Please implement a command line REPL (read-eval-print loop) that drives a simple in-memory key/value storage system. This system should also allow for multi-level nested transactions. A transaction can then be committed or aborted.

Please avoid spending more than 4 hours on this exercise.

## Example Run
```
$ my-program
> WRITE a hello
> READ a
hello
> START
> WRITE a hello-again
> READ a
hello-again
> START
> DELETE a
> READ a
Key not found: a
> COMMIT
> READ a
Key not found: a
> WRITE a once-more
> READ a
once-more
> ABORT
> READ a
hello
> QUIT
Exiting...
```

## Command-line arguments
    READ   <key> Reads and prints, to stdout, the val associated with key. If
           the value is not present an error is printed to stderr.
    WRITE  <key> <val> Stores val in key.
    DELETE <key> Removes all key from store. Future READ commands on that key
           will return an error.
    START  Start a transaction.
    COMMIT Commit a transaction. All actions in the current transaction are
           committed to the parent transaction or the root store. If there is no
           current transaction an error is output to stderr.
    ABORT  Abort a transaction. All actions in the current transaction are discarded.
    QUIT   Exit the REPL cleanly. A message to stderr may be output.

## Other Details
* For simplicity, all keys and values are simple ASCII strings delimited by whitespace. No quoting is needed.
* All errors are output to stderr.
