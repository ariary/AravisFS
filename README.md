# CeasarFS

A fake encrypted file system 🌺 *Another non-production ready software*

	🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟
**DRAFT STAGE - *Any idea, criticism, contribution is welcome***
No pretension just to learn and keep my mind busy

	🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟
## 🔦 Idea
Providing a fake encrypted FS and utilities to interact with

**Why fake?**

 - Cause encryption isn't strong
 - Cause it does not provide a real fs, just a representation of it. *(And to be honest it only encrypts a folder but by extension we say a file system)*

### Use case
 - To avoid your volume being spied on by your cloud provider
 - To make a temporary folder on the target machine if you are a black hat and you do not want to be spied on by forensic people. Or if you want to hide a payload.
 - To boast of having an encrypted fs .. but unusable

## 💺 Usage (under construction)

### Encrypt a fs

 - **Encrypt with `CeasarSalad`**
 *parameter*: `folder`
 *return*: `ceasar_salad`
 *process*: Take `folder`, encrypt it using magic, and return the representation of the folder encrypted `ceasar_salad`

### Decrypt an encrypted fs
 - **Decrypt with `CeasarDalas`**
 *parameter*: `ceasar_salad`
 *return*: `folder_representation`
 *process*: Take `ceasar_salad`, decrypt it using magic, and return the representation of the folder decrypted `folder_representation`

 - **Create the fs described by a  representation with `CeasarTexas`**
 *parameter*: `folder_representation` (representation of an unencrypted folder/fs)
 *process*: Create the different folder and file describe in the `folder_representation` on the file system.
 
 #### Example
 First I encrypt a fs:

     ceasarsalad my_folder > secret_fs

 
 Then I put `secret_fs`on my USB key and want to use it on another PC (why??):

    $ ceasardalas secret_fs > unencrypted_fs_representation
    $ ceasartexas unencrypted_fs_representation

### Explore an encrypted fs
 - **Browse encrypted fs with  `ceasarls`**
 *parameter*: `folder_representation` (representation of an encrypted fs) and `path`
 *process*: Same thing that `ls path`
 
  - **Print the content of a file of the encrypted fs with  `ceasarcat`**
 *parameter*: `folder_representation` (representation ofan encrypted fs) and `filepath`
 *process*: Same thing that `cat filepath`
 *Pb: leak  info in stdout that could be spied*
 #### Example
 On a target machine containing my encrypted fs `secret_fs`, Iwan to know the content of the folder toto:

     (on my machine) $ ceasarpath toto 
     ~~> **copy/paste result**
     (on the target machine) $ ceasarls secret_fs <paste_ceasarpath _result>

#### Note
CeasarPath will be provided later. The reason why we encrypt the path is to avoid info leak from `.bash_history` (information about the hierarchy of the secret fs)

### Manipulate an encrypted fs (future)
`ceasarmv`,  `ceasarcp`, `ceasarrm`,  `ceasartouch`, etc

