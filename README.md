## gossh 
*Developing*

This is not ssh. This is a help tool for ssh.

### gossh can de what?
1. quickly connect to a host(of course you donot have to type password), and then do shell commands or copy sths.
2. support alias of hostname(very usefull if your hostname is very long)

gossh has two parts, server and client.

### How to use gossh
1. You need to register a account of [server](doc/server.md).
2. Add `hostname, username, password` to server. server will store it encrypted.

I suppose you already finished the first two steps.
And record of `example-vps.com, root, 123456` has been added.

	# client use such command to connect to example-vps.com
	# usally we use ssh root@example-vps.com echo hi, and then password has to input.
	# but now with gossh, we can use
	#
	gossh root@example echo hi

### The working flow.
1. client login to server, server check auth
2. client send user(root), hostAlias(example) to cloud server
3. server search hostname which user=root and hostname matches example.
4. server send back avalible list of (hostname, password)
5. client call ssh, scp or rsync to do commands (use sshpass of simply gosshpass)

### Other ideas
for a group use, give them a group-account.
for personal use, give them a single-account and group is group-account

so user can use themself conf and also the group conf.

### How to join
All contributes can be found in [todo](doc/todo.md)

