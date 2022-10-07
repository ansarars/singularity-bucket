Singularity plugin for AWS S3 buckets
=====================================

This directory contains an example CLI plugin for singularity. It
demonstrates how to add a command and flags.

Pre-requisites for Singularity plugin
------------------------------------
    1. Install system dependencies
        On Debian-based systems, including Ubuntu:
        
        # Ensure repositories are up-to-date
        sudo apt-get update
        
        # Install debian packages for dependencies
        sudo apt-get install -y \
        build-essential \
        libseccomp-dev \
        pkg-config \
        squashfs-tools \
        cryptsetup

    2. Install Go version 1.19
        $ export VERSION=1.19 OS=linux ARCH=amd64 && \  # Replace the values as needed
        wget https://dl.google.com/go/go$VERSION.$OS-$ARCH.tar.gz && \ # Downloads the required Go package
        sudo tar -C /usr/local -xzvf go$VERSION.$OS-$ARCH.tar.gz && \ # Extracts the archive
        rm go$VERSION.$OS-$ARCH.tar.gz    # Deletes the ``tar`` file
        
        $ echo 'export PATH=/usr/local/go/bin:$PATH' >> ~/.bashrc && \
        source ~/.bashrc
        
        $ go version
        go version go1.19 linux/amd64

    3. Install Singularity version 3.9.6
        $ git clone https://github.com/sylabs/singularity.git
        
        $ cd singularity
        
        $ git checkout v3.9.6
        
        $ ./mconfig && \
        make -C builddir && \
        sudo make -C builddir install
        
        $ singularity version
        3.9.6

    4. Install AWS cli
        $ curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
        $ unzip awscliv2.zip
        $ sudo ./aws/install
        # Verify aws version
        $ aws --version
        aws-cli/2.8.1 Python/3.9.11 Linux/4.15.0-20-generic exe/x86_64.ubuntu.18 prompt/off
    
    5. Create AWS account on https://aws.amazon.com/ and refer below link for generating
       aws_access_key_id and aws_secret_access_key using AWS console.

       https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html

    6. Configure AWS credentials on the system
        $aws configure
        AWS Access Key ID [None]: XXXXXXXXZZZZZXX8YYXX
        AWS Secret Access Key [None]: xxxxxt7qRc1nqmOxxxxlHKnJ0yPmeHxxxxxxxxxx
        Default region name [None]: us-west-2
        Default output format [None]: json

        # After configuring AWS credentials verify that aws_access_key_id and aws_secret_access_key values are there
        # in ~/.aws/credentials file
        $ cat ~/.aws/credentials
        [default]
        aws_access_key_id = XXXXXXXXZZZZZXX8YYXX
        aws_secret_access_key = xxxxxt7qRc1nqmOxxxxlHKnJ0yPmeHxxxxxxxxxx


Building singularity plugin
---------------------------
Obtain a copy of the source code by running:

    git clone https://github.com/ansarars/singularity-s3-bucket.git
    cd singularity-s3-bucket

Still from within that directory, run:

	singularity plugin compile .

This will produce a file `singularity-s3-bucket.sif`.


Install singularity plugin from SIF file
----------------------------------------

Once you have compiled the plugin into a SIF file, you can install it
using following command:

	$ singularity plugin install singularity-s3-bucket.sif

Singularity will automatically load the plugin code from now on.

Other commands
--------------

You can query the list of installed plugins:

    $ singularity plugin list
    ENABLED  NAME
        yes  singularity-s3-bucket-plugin

Disable an installed plugin:

    $ singularity plugin disable singularity-s3-bucket-plugin

Enable a disabled plugin:

    $ singularity plugin enable singularity-s3-bucket-plugin

Uninstall an installed plugin:

    $ singularity plugin uninstall singularity-s3-bucket-plugin

And inspect a SIF file before installing:

    $ singularity plugin inspect singularity-s3-bucket.sif
    Name: singularity-s3-bucket-plugin
    Description: singularity plugin for s3 bucket operations
    Author: Arshad Alam Ansari
    Version: v1.0


Bucket Usage:
------------
Following command would display the usage of the bucket command:

    $ singularity bucket --help

    bucket

    Usage:
    singularity bucket <create|list> <name>


    Options
    - create
        Specifies bucket creation operation.
    - list
        Specifies bucket list operation.
    - name
        Specifies bucket name. Required for create operation only.


    Description:Allows life-cycle management of a bucket

    Options:
    -h, --help   help for bucket


    Examples:singularity bucket list


    For additional help or support, please visit https://www.sylabs.io/docs/




bucket operation examples:
-------------------------
1.Create bucket:

    Command:
    $ singularity bucket create name=first-s3-bucket-fff


2.List bucket:

    Command:
    $ singularity bucket list

    Response:
    Found bucket: first-s3-bucket-eee, created at: 2022-10-07 06:23:47 +0000 UTC
    Found bucket: first-s3-bucket-fff, created at: 2022-10-07 14:41:19 +0000 UTC


Object Usage:
------------------------
Following command would display the usage of the object command:

    $ singularity object --help

    object

    Usage:
    singularity object <list> <bucket-name>


    Options
    -list
        Specifies bucket list operation.
    -bucket-name
        Specifies bucket name. Required for object list operation.


    Description:Allows life-cycle management of a objects

    Options:
        -h, --help   help for object


    Examples:singularity object list name=bucketName


    For additional help or support, please visit https://www.sylabs.io/docs/



object operation examples:
-------------------------------------
1.list bucket objects:
    
    Command:
    $ singularity object list bucket-name=first-s3-bucket-ccc

    Response:
    Name: aaa.txt, Last Modified: 2022-10-07 00:00:46 +0000 UTC
    Name: bbb.txt, Last Modified: 2022-10-07 00:02:23 +0000 UTC


