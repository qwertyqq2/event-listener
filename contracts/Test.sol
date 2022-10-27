// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


contract Test {
   event CreateDeposit(address _from, uint256 _value);
   event CloseDeposit(address _from);

   mapping(address=>uint) public deposites;

   function Deposit() public payable {      
      emit CreateDeposit(msg.sender, msg.value);
      deposites[msg.sender] = msg.value;
   }

   function Close() public{
      require(deposites[msg.sender]!=0, "Not exist");
      emit CloseDeposit(msg.sender);
      deposites[msg.sender] = 0;
   }


}