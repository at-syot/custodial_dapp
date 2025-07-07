// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract Todo {
  Task[] tasks;
  struct Task {
    string content;
    bool status;
  }

  constructor() {
  }

  function add(string memory content) public {
    tasks.push(Task(content, false));
  }

  function get(uint id) 
    public 
    view
    returns (Task memory) {
      return tasks[id];
  }

  function list() public view returns (Task[] memory) {
    return tasks; 
  }

  function update(uint id, string memory content, bool status) 
    public {
      tasks[id].content = content;
      tasks[id].status = status;
  }

  function remove(uint id) public {
    delete tasks[id];
  }
}
