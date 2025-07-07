// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.0;

contract Ballot {
  struct Voter {
    uint weight; 
    bool voted;
    address delegate;
    uint vote;
  }

  struct Proposal {
    bytes32 name;
    uint voteCount;
  }

  address public chairperson;
  mapping(address => Voter) public voters;
  Proposal[] public proposals;

  constructor(bytes32[] memory proposalNames) {
    chairperson = msg.sender;
    voters[chairperson].weight = 1;

    for (uint i = 0; i < proposalNames.length; i++) {
      proposals.push(
        Proposal({ name: proposalNames[i], voteCount: 0})
      );
    }
  }

  function giveRightToVote(address voter) external {
    require(msg.sender == chairperson, "only chairperson can give the right to vote.");
    require(!voters[voter].voted, "The voter already voted.");
    require(voters[voter].weight == 0, "No need to give right to voter");
    voters[voter].weight = 1;
  }

  function delegate(address to) external {
    Voter storage sender = voters[msg.sender];
  
    require(sender.weight != 0, "sender:You not allow to vote");
    require(!sender.voted, "sender:You already voted");
    require(msg.sender != to, "Self-delegation is not allow");

    // it try to find a voter that isn't delegate their vote to any address
    // and the untimate goal is to find the terminal voter that ***we will give our voting power to.
    while (voters[to].delegate != address(0)) {
      to = voters[to].delegate;

      require(to != msg.sender, "We found a delegation loop");
    }

    Voter storage _delegating = voters[to];
    require(_delegating.weight >= 1, "delagated voter:This voter have no right to vote");

    sender.voted = true;
    sender.delegate = to;

    // this is another functionality
    if (_delegating.voted) {
      // if delegated voter already voted,
      // immediatly give their vote's weight to the proposal
      proposals[_delegating.vote].voteCount += sender.weight;
    } else {
      _delegating.weight += sender.weight;
    }
  }

  function vote(uint proposal) public {
    Voter storage voter = voters[msg.sender];
    require(voter.weight >= 1, "You don't have voting right.");
    require(!voter.voted, "You already voted.");

    voter.voted = true;
    voter.vote = proposal;

    proposals[proposal].voteCount += voter.weight;
  }

  /// @dev Computes the winning proposal taking all
  /// previous votes into account.
  function _winningProposal() 
  internal view
  returns (uint winningProposal_){
    uint winningVoteCount = 0;
    for(uint p = 0; p < proposals.length; p++) {
      if (proposals[p].voteCount > winningVoteCount) {
        winningVoteCount = proposals[p].voteCount;
        winningProposal_ = p;
      }
    }
  }

  function winningName() 
  public view
  returns (bytes32 winningName_) {
    winningName_ = proposals[_winningProposal()].name;
  }
}
