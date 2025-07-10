// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BlindAuction {
    struct Bid {
        bytes32 blindedBid;
        uint deposit;
    }

    address public highestBidder;
    uint public highestBid;

    mapping(address => uint) public pendingReturns;
    mapping(address => Bid[]) public bids;

    event HighestBidIncreased(uint highestBid);
    event AuctionEnded(address winner, uint highestBid);

    error BidNotHighEnough(uint highestBid);

    constructor() {}

    function bid(bytes32 blindedBid) external payable {
        Bid memory _bid = Bid({blindedBid: blindedBid, deposit: msg.value});
        bids[msg.sender].push(_bid);
    }

    function reveal(
        uint[] calldata values,
        bool[] calldata fakes,
        bytes32[] calldata secrets
    ) external {
        uint bidsLen = bids[msg.sender].length;
        require(values.length == bidsLen);
        require(fakes.length == bidsLen);
        require(secrets.length == bidsLen);

        uint refund;
        for (uint i = 0; i < bidsLen; i++) {
            (uint value, bool fake, bytes32 secret) = (
                values[i],
                fakes[i],
                secrets[i]
            );
            Bid storage checkingBid = bids[msg.sender][i];
            if (
                checkingBid.blindedBid !=
                keccak256(abi.encodePacked(value, fake, secret))
            ) {
                continue;
            }

            refund += checkingBid.deposit;
            if (!fake && checkingBid.deposit >= value) {
                if (placeBid(msg.sender, value)) refund -= value;
            }
            checkingBid.blindedBid = bytes32(0);
            checkingBid.deposit = 0;
        }
        payable(msg.sender).transfer(refund);
    }

    // So!, placeBid fn is not just place to bid of that bidder
    // but also deteminds protential newly highest bidder and
    // when the highest bid is incresed the
    // old bidder's bid will store to pendingRefund(prevHighestBidder) = prevHighestBid
    function placeBid(address bidder, uint value) internal returns (bool) {
        if (value <= highestBid) {
            return false;
        }

        address previousBidder = highestBidder;
        uint previousHighestBid = highestBid;
        if (previousBidder != address(0)) {
            pendingReturns[previousBidder] += previousHighestBid;
        }

        highestBidder = bidder;
        highestBid = value;
        emit HighestBidIncreased(highestBid);

        return true;
    }

    function withdraw() external payable returns (bool) {
        uint amount = pendingReturns[msg.sender];
        if (amount > 0) {
            // Check -> Effect -> Interact pattern
            pendingReturns[msg.sender] = 0;

            address payable _sender = payable(msg.sender);
            if (!_sender.send(amount)) {
                pendingReturns[msg.sender] = amount;
                return false;
            }
        }

        return true;
    }
}
