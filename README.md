

# AravisFS 🗻🌄

A fake encrypted file system 🔐 *Another non production-ready software*

	🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟

| **DRAFT STAGE** - *Any idea, criticism, contribution is welcome*.  |
|:------------------------------------------------------------------------------------------------------------------:|
*No pretension just to learn and keep my mind busy*
	🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟
	
	
## 🔦 Idea
**Aim?**
Providing a fake encrypted FS and utilities to interact with. 

Our objective is to leak as little information as possible. The local machine is our trusted environnement where we could manipulate the key, the clear-text data. This is our **(light side)**
 And on another untrusted location we have our encrypted fs. We do not want to manipulate key or clear text but we want to be able to interact as much as possible with the encrypted fs. This is our **(dark side)**

For this purposewe use 2 utilities, each on different side:
- `adret`: Encrypt/decrypt fs. Deal with key & clear-text data ***(light side)***
 - `ubac`: Interact with encrypted fs. Deal with no "sensitive" data (no key & clear-text manipulation) ***(dark side)***

*We accept to leak information about the fs structure (number of file/folder, size)*



**Use case?**
 - To avoid your volume being spied on by your cloud provider
 - To make a temporary folder/fs on a target machine if you are a black hat and you do not want to be spied on by forensic people. Or if you want to hide a payload.
 - To boast of having an encrypted fs .. but practically unusable anyway

**Why "fake"?**
 - Cause encryption isn't strong
 - Cause it does not provide a real fs, just a representation of it. *(And to be honest it only encrypts a folder but by extension we say a file system)*

 
 ## 💺 Usage (***In case I ever code***)
### 🔍 Explore encrypted folder
#### List folder content from my encrypted fs
First I encrypt my fs :

    (local) $ adret -encrypt_fs <myfolder> <mykey>

Then I put the result (encrypted fs `.arafs`) and `ubac` utility on the dark zone (target machine, my OVH server, container on GKE, etc) the way I want. I could then remove `<myfolder>`from my host, otherwise it has real no sense

Say I want to `ls` in `<myfolder>/toto/`, so I encrypt the path:

    (local) $ adret -encrypt_path <myfolder>/toto/ <mykey>
    **We obtain darkpath! COPY THE RESULT**

Then we do the `ls` in our encrypted fs:

    (remote) ubac -ls darkpath myencryptedfs.arafs
    **We obtain ls result! COPY THE RESULT**

And finally decrypt the result of ls:

    (local) $ adret -encrypt_path <ls_rult> <mykey>

#### Print file content from my encrypted fs
Idem as above but change the `ubac command with:

    (remote) ubac -cat darkpath myencryptedfs.arafs

### 🤖 Automatize a bit 
If you want to interact with your remote encrypted fs more fluidly

**Prerequisites**

 - my encrypted fs + `ubac`on remote
 - remote is accesible on port `<ip>:<port>`
 - `adret` on local 


Start your `ubac` listener in the same location where the encrypted fs is:

    (remote) $ ubac -listen <port>

Connect to the listener with your local machine:

    (local) $ adret -connect <ip>:<port>

I can now interact directly with my encrypted fs on my local machine and obtain direct result:

    (remote) $ adret -remotecat /random/topath/file

### ✂️ Manipulate encrypted fs
#### Remove a file from encrypted fs
    (remote) ubac -rm darkpath myencryptedfs.arafs
 
Use `-mv`, `-touch` etc the same way you could use it in unix system

## 📝 Spec
### Target 1 - MVP

#### 🌄 Adret
| function    | parameter | return         | use                                                                 
|-------------|-----------|----------------|---|
| encrypt     | key, data | encrypted_data | Encrypt/Decrypt  `data`  |
| encrypt_fs  | key, path | encrypted_fs   | Encrypt fs (=folder pointed by the `path`)  |
| darken_path | key, path | encrypted_path | Encrypt a `path`                                            |

`encrypted_fs`is a file (`.arafs`) representing the folder hierarchy (~fs) encrypted. ***TO DO*** Determine the structure (tree, yaml etc )

#### 🗻 Ubac
| function | parameter                   | return               | use                        |
|----------|-----------------------------|----------------------|----------------------------|
| ls       | encrypted_fs, darkened_path | encrypted_result     | Get encrytped `ls` result  |
| cat      | encrypted_fs, darkened_path | encrypt_content_file | Get encrypted `cat` result |

### Target 2 - The future is now
#### 🌄 
| function | parameter | return | use                                                                  |
|----------|-----------|--------|----------------------------------------------------------------------|
| remotels, remotecat, etc   | host, port, path, key     |        | connt with `ubac`  listener to perform different function on the encrypted fs |

#### 🗻
| function | parameter | return | use                                                                  |
|----------|-----------|--------|----------------------------------------------------------------------|
| listen   | port      |        | Act like a server. Wait for request from ubac, process it, return it |

### target 3 - The world is yours
#### 🗻
| function | parameter                   | return | use                       |
|----------|-----------------------------|--------|---------------------------|
| mv       | encrypted_fs, darkened_path |        | mv in the encrypted fs    |
| cp       | encrypted_fs, darkened_path |        | cp in the encrypted fs    |
| rm       | encrypted_fs, darkened_path |        | rm in the encrypted fs    |
| touch    | encrypted_fs, darkened_path |        | touch in the encrypted fs |

And theirs siblings `remotemv`  ,etc

## 🧙 How does it works?
Magic! (**soon explained**)

### How  is the fs encrypted ?

When the `adret`utlity encrypt a folder it returns what we called the **"encrypted filesystem"** wich is `.arafs`. It is in fact a text file of the fs representation (encrypted of course).

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
Take the example of `ubac -ls darkpath myencryptedfs.arafs` which aims to perform `ls` of the `darkpath` in `myencryptedfs.arafs`.

First it will search the `darkpath` (it is an encrypted path) name in the list which is `myencryptedfs.arafs` . If the type is file it will return ... If the type is a folder it will take the content part and retrieve the type of all resource within (by searching with their name).
Finally it will return a list with the child resources name with their type. (child resource name are encrypted)

### How does the remote interaction work  with the encrypted fs?
Here the sequence of a `remotecat`. 

Previously we have our `ubac` listener launch on an accesible port on remote and the encrypted fs in the same location as the listener:

 1. `-remotecat path/to/file <key>`
 2. The path is encrypted using the `darken_path`function ~~> `darkenedpath`
 3. The adret send a request to the ubac listener to perform a `cat` with `darkenedpath`
 4. The ubec listener perform the cat and return the encrypted result to `adret`
 5. `adret` decrypt it and print the result

