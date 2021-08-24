
<h1 align="center"> AravisFS 🗻🌄</h1>
<p align="center">
	A remote fake encrypted filesystem  🔐 <i>Another non production-ready software</i>
</p>

<div align="center">
	🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟

	Any idea, criticism, contribution is welcome
	No pretension just to learn and keep my mind busy
</div>
	
----

<p align="center">
<strong><a href="#-idea">🔦 Idea</a></strong>
|
<strong><a href="#-installation">💺 Installation</a></strong>
|
<strong><a href="#-usage">🚀 Usage</a></strong>
|
<strong><a href="#-spec">📝 Spec</a></strong>
|
<strong><a href="#-how-does-it-work">🧙 How does it work?</a></strong>
|
<strong><a href="#limitsimprovements">💭Limits/improvements</a></strong>
</p>

----

	
## 🔦 Idea
**Aim?**
Providing a fake encrypted FS and utilities to interact with.  The objective is to leak as little information as possible. 

The local machine is our trusted environment where we could manipulate the key, the clear-text data etc. This is our **(light side)**
 On another untrusted location we have our "encrypted file system". We do not want to manipulate key or clear text but we want to be able to interact as much as possible with the encrypted fs. This is our **(dark side)**

For this purpose we use 2 utilities, each on different side:
- `adret`: Encrypt/decrypt fs etc. Deal with key & clear-text data ***(light side)***
 - `ubac`: Interact with encrypted fs. Deal with no "sensitive" data (no key & clear-text manipulation) ***(dark side)***

*We accept to leak information about the fs structure (number of file/folder, size)* on the dark side

***Note:*** `adretctl` offer a much clever user experience of adret utility with a CLI etc. See [usage](#-usage)



**Use case?**
 - To avoid your volume being spied on by your cloud provider
 - To make a temporary folder/fs on a target machine if you are a black hat and you do not want to be spied on by forensic people. Or if you want to hide a payload.
 - To get a manipulable ransomware 
 - To boast of having an encrypted fs .. but practically unusable anyway

**Why "fake"?**
 - Cause encryption isn't strong enough *(AES ECB)*
 - Cause it does not provide a real fs, just a representation of it. *(And to be honest it only encrypts a folder but by extension we say a filesystem)*

## 💺 Installation

 Clone the repo and download the dependencies locally:
```    
git clone https://github.com/ariary/AravisFS.git
make before.build
```

 To build `adret` :

     make build.adret
    
 To build `adretctl` :

     make build.adretclt
	    
Idem, to build `ubac`:

    make build.ubac

***REMINDER***: use `adret`/`adretctl` in an trusted environment cause it will manipulate clear-text data and key. Transfer `ubac` utility where you encrypted fs is (w/ tftp, python web server, nc, certutil, etc Be [creative](https://medium.com/@PenTest_duck/almost-all-the-ways-to-file-transfer-1bd6bf710d65))

## 🚀 Usage 
### Adretctl & ubac demo
 ![demo](https://github.com/ariary/AravisFS/blob/main/img/adretctldemo.gif)

*In this demo, I have an encrypted fs (encrypt with key "toto") on the non trusted zone. `ubac listen` expose it and interact with it using `adretctl` : browse it, cat file and delete a directory*

See [adretctl spec](#target-3---to-the-infinite---interactive-prompt-w-fs-cli)

### 🔒 Encrypt folder/fs

See [demo encrypt (.GIF)](https://github.com/ariary/AravisFS/blob/main/img/encryptfs.gif)

### 🔍 Explore encrypted folder
<details>
	<summary>List folder content from my encrypted fs</summary>
First I encrypt my fs :

    (local) $ adret encryptfs -key=<secret> <path>
    [...]
    test/mytestfolder/titi
    [...]
    done! Encrypted fs saved in encrypted.arafs

Then I put the result, our encrypted fs `.arafs`, and `ubac` utility on the dark zone the way I want( It  could be a target machine, my OVH server, container on GKE, etc).  I could then remove `<myfolder>`from my host, *otherwise it has real no sense* (see [example](https://github.com/ariary/AravisFS/blob/main/img/encryptfs.gif))

Say I want to `ls` in ` "test/mytestfolder/titi"`, so I encrypt first the encrypt the path:

    (local) $ adret darkenpath -key="toto" "test/mytestfolder/titi"
    **We obtain encrypted path! COPY THE RESULT**

Then we do the `ls` in our encrypted fs:

    (remote) ubac ls -path=myencryptedfs.arafs <encryptedpath_from_adret> 
    **We obtain ls result! COPY THE RESULT**

And finally decrypt the result of ls:

    (local) $ adret decryptls -key="toto"
    tutu tutu.txt utut.txt
</details>

<details>
	<summary>Print file content from my encrypted fs</summary>
Idem as above with `ls` but change the `ubac` command with:

    (remote) ubac cat -path=myencryptedfs.arafs <encryptedpath_from_adret>
</details>
<details>
	<summary>Print  encrypted fs tree</summary>
First retrieve encrypted tree from encrypted fs:

    (remote) ubac tree test/arafs/medium.arafs
    **We obtain the encrypted tree result! COPY THE RESULT**
   
  Then decrypt it to print it with (assume the fs was encrypted with the key "toto"):

      (local) $ adret decrypttree -key="toto" <encryptedtree_from_ubac>
</details>
   

### 🤖 Automate a bit 
If you want to interact with your remote encrypted fs more fluidly

**Prerequisites**

 - my encrypted fs + `ubac`on remote
 - remote is accesible on port `<ip>:<port>`
 - `adret` on local 
 - have the `key` which encrypt the fs

<details>
<summary>List folder content from remote encrypted fs</summary>
Start your `ubac` listener on the remote where the encrypted fs is :

    (remote) $ ubac listen -path="./test/arafs/encrypted.arafs" 4444
	** IT WILL LAUNCH AN HTTP SERVER WAITING FOR REQUEST ON PORT 4444***

An local machine configure your environment variable to dial with remote ubac listenerConnect to the listener with your local machine:

    (local) $ eval `adret configremote -port="4444" -host=<remote_ubac_hostname>`

I can now interact directly with my encrypted fs on my local machine and obtain direct result:

    (remote) $ adret remotels -key=toto test/mytestfolder
</details>

<details>
	<summary>List folder content from remote encrypted fs</summary>
Launch ` ubac` listener and config local host wit `adret` like [above](#list-folder-content-from-remote-encrypted-fs) and launch `remotecat`:

    (remote) $ adret remotecat -key=toto test/mytestfolder/toto.txt
</details>
	
### ✂️ Manipulate encrypted fs
<details>
	<summary>Remove a file from encrypted fs</summary>

First retrieve encrypted tree from encrypted fs:

    (remote) ubac tree test/arafs/medium.arafs
    **We obtain the encrypted tree result! COPY THE RESULT**

Then get a patch (which describes the changes you want to make):
```
(local) adret rmpatch -key=toto -tree=<output_ubac_tree> <resource_I_want_to_remove>
***WE OBTAIN A PATCH COPY IT***
```

Then apply the patch onto the encrypted fs
```
(remote) ubac applypatch -path=encrypted.arafs <my_patch>
``` 
</details>
	
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




## 🧙 How does it work?
Magic!

### How  is the fs encrypted ?

When the `adret`utlity encrypt a folder it returns what we called the **"encrypted filesystem"** wich has `.arafs` extension. It is in fact a text file of the fs representation (encrypted of course).

The structure of  the content of the `.arafs` file  is influenced by the following constraint: **it musts be used to retrieve data without using the key**

#### Encrypted fs data structure
For this reason, the data stracture of the content is a (unordered) list of all the **"resources"**  present in the folder. In that way, we could only know how many resources the encrypted fs has and also the size of the fs without having the key.

##### Resource
A resource is:
|  |  ||
|--|--|--|
|🔐name  | type |🔐content|

**name**: is the encrypted value of the path of the resource
**type** : is the type of the ressource within the fs. it could be:
- folder
-  file

**content**: is the content of the resource ie a list of child resources name if type is  folder or the file content if type is file

#### Example: Search in encrypted fs

Take the example of  `ubac ls -path=myencryptedfs.arafs darkpath` which aims to perform `ls` of the `darkpath` in `myencryptedfs.arafs`.

First it will search the `darkpath` (it is an encrypted path) name in the list which is `myencryptedfs.arafs` . 
If the type is file it will return the filename. 
If the type is a folder it will take the content part which is all its child resources encyrpted.

### How does the remote interaction work  with the encrypted fs?
Here the sequence of a `remotecat`. 

Previously we have our `ubac` listener launch on an accesible port on remote and the encrypted fs in the same location as the listener:

 1. `remotecat  -key=<key> path/to/file`
 2. The path is encrypted using the `darkenpath`function ~~> `darkenedpath`
 3. The adret send a request to the ubac listener to perform a `cat` with `darkenedpath`
 4. The ubec listener perform the cat and return the encrypted result to `adret`
 5. `adret` decrypt it and print the result

### How does the fs is modified ?

As the encrypted FS is represented in a JSON file format and `ubac` has no acknowledge about what is inside (and couldn't obtain), we must have 5 steps to modify the encrypted fs 
 1. Ask `ubac` to get the tree of the encrypted FS
 2. `Darkenpath` the parent directory of the resource which will be modified (or added)
 3. Ask `ubac` content of the parent directory (to modify it also)
   * if the resource is a directory, it will use the tree to delete all resources under it (which the Tree structure we won't have to launch it recursively)
 5. With the tree, craft the patch to apply on the FS with `adret`. Patch depend of the logic (if you want to remove, add a ressource etc)
 6. Provide the patch to `ubac` to apply it on the FS

##### Tree
Tree is a containaing all the Resource in the ecrypted fs. It is a `Node` list

A `Node` on ubac side is:
|  |  |
|--|--|
|🔐name  | type|

On adret side, it is the same concept. A tree is a list of Node but the `Node` contains another field `Parent` which is the resource parent directory ( the resource is represented by `name` field in Node):
|  |  |	 |
|--|--|--|
|🔐name  | type |🔐Parent|

where `Parent` is the prefix of the resource name (ie resource parent directory).

##### Patch
Patch is used to inform `ubac` which resources it will change. So it contains 3 list: `to_delete` for resources that must be removed, `to_add` for resources that must be added, `to_change` for resources that must be change (theirs contents).

(`to_add` & `to_change` contain a list of the resource with their content associated)

#### Example: remove an element

 1. ask ubac the tree
 2. Ask adret the patch to delete the resource
  add to `to_delete` the resource + modify parent resource content (withdraw the resource of it) and add parent resource to `to_change`
	 - if resource is a folder: add to `to_delete` all the resource with prefix containing <resource_to_delete_name> (ie under the directory)
 3. Provide the the patch to ubac, and let ubac apply it

## 💭Limits/improvements
- `adret` encrypt using AES ECB (attack possible). *A more robust AES encryption scheme change the nonce at each encryption => for the same input different outputs at each encryption. It is a problem as darkenpath must provide the same path encrypted as the initial encryption (when we encrypt the fs)*
- you can't encrypt a filesystem w/ a folder/file having `'\'` in its name. *It is due to the way we encapsulate resource in directory content*
- you can't perform "`ls .`". *As we search for resource with its exact name we do not have `.` resource from now*
- launch `adret` in the same directory of the fs you want to encrypt *Otherwise it will keep prefix like "../" etc, (see filesystem.GetDirectoryContent)*
- for `adret decrypt cat` we could not `cat` big file. *It is due to the fact that we take the encrypted content to show from command-line. Hence we are limited by `ARG_MAX` length (`getconf ARG_MAX`to know the value). For such reason avoid embedding binary file (or try to compress it using `upx` command before). This limitation applies for all command in fact*
	- Workaround for `ARG_MAX` length: Save the arg in file and pass it with `$(cat <FILE>)`
- file permissions
- CLi flag and command parsing is homemade (see [cobra](https://github.com/spf13/cobra) to improve/facilitate)

