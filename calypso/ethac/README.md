Ethereum Access Control


In order to run this you will need to be running a version of the ganache client
running on your localhost to be able to test it. It is running on port 7545 in the
connectionString in the code. Many of the values that are used for testing are 
hardcoded from the ethereum client so if you want to try a custom configuration
you will have to change some of the code. If you want to run the same starting 
configuration that was done while writing this code you must do the following.

Run the CI.sh file. It should set up the necessary environment variables for you 
to be able to run all of the commands. Here are the things that it will do for you.

It makes the assumption that you have go setup on your computer and that your gopath is the default $HOME/go/....

Run the go get ./... command

-Install node.js. You'll require node.js to run the tests for deploying and interacting
with deployed contracts on the ethereum with truffle and run truffle

-install the necessary node packages>
    - truffle
    - ganache-cli
    - web3
        - beta version, current head is not stable
    - mocha

Install the Solidity compiler

Setup the abigen tool from the go ethereum so you can build go files from solidity smart contracts

Build the Calypso.go contract from the Calypso.sol file using the abigen tool.
