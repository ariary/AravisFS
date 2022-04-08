## 📝 Spec
### Target 1 - MVP - Basic utility on FS

#### 🌄 Adret
| function    | parameter | return         | use                                                                 
|-------------|-----------|----------------|---|
| darkenpath     | key, path | encrypted_path | Encrypt a path   |
| encryptfs  | key, path | encrypted_fs   | Encrypt fs (=folder pointed by the path)  |
| decryptls | key, path |  |  decrypt `ls` result (from ubac utility)
| decryptcat | key, path |  |  decrypt `cat` result (from ubac utility)
| decrypttree | key, resources_tree |  |  decrypt tree result (from ubac utility)                                               |

`encrypted_fs`is a file (`.arafs`) representing the folder hierarchy (~fs) encrypted. 

#### 🗻 Ubac
| function | parameter                   | return               | use                        |
|----------|-----------------------------|----------------------|----------------------------|
| ls       | encrypted_fs, darkened_path | encrypted_result     | Get encrytped `ls` result  |
| cat      | encrypted_fs, darkened_path | encrypt_content_file | Get encrypted `cat` result |
| tree      | encrypted_fs | encrypt_resources_tree | Get encrypted `tree` result (malformed ie in json) |

### Target 2 - The future is now - Remote interaction with FS

#### 🌄 
| function | parameter | return | use                                                                  |
|----------|-----------|--------|----------------------------------------------------------------------|
| configremote   | host, port     |   cmdline_setting_envvar     | return command line to set env var which will be used to dial with remote ubac |
| remotels, remotecat, etc   | resource_name, key    |    remote_command_result_decrypted    | connect with `ubac`  listener to perform different function on the encrypted fs and print the decrypted result|

#### 🗻
| function | parameter | return | use                                                                  |
|----------|-----------|--------|----------------------------------------------------------------------|
| listen   | port      |        | Act like a server. Wait for request from ubac, process it, return it |

### Target 3 - to the infinite - Interactive prompt w/ FS (CLI)
Build a interactive prompt CLI for adret: **`adretctl`**. It is used to dial with an `ubac` listener in an interactive way.

it must reimplemet the already present function (`ls`, `cat`,`tree`) and add some more context to browse the the remote encrypted filesystem (`cd`).
| function |  use                                                                  |
|----------|----------------------------------------------------------------------|
| keyconfig| set key use to decrypt/encrypt| 
| keyprint| print the current key used| 
| connect|connect to remote ubac listener|
| cat,rm,cd,tree,ls| as we expect|  

### Target 4 - The world is yours - Manipulate FS


#### 🌄 
| function | parameter                   | return | use                       |
|----------|-----------------------------|--------|---------------------------|
| rmpatch       | key,tree,resource|        | provide patch that remove the resource in the fs (represented by the tree)   |

the `tree` is obtained by asking it to `ubac` first

(probably mvpatch, cppatch, vimpatch)

#### 🗻
| function | parameter                   | return | use                       |
|----------|-----------------------------|--------|---------------------------|
| applypatch| encrypted_fs, patch|        | apply a patch onto the encrypted fs    |

