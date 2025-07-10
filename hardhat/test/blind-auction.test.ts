import { expect, assert } from "chai";
import { ethers } from "hardhat";
import { BigNumberish, BytesLike, toBigInt } from "ethers";
import type { BlindAuction } from "../typechain-types";
import type { HighestBidIncreasedEvent } from "../typechain-types/BlindAuction";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";

type BidCommit = {
  deposit: BigNumberish; // exclude in hash fn
  secret: BytesLike;
  fake: boolean;
};

// Used for - reveal function
type RevealArgs = {
  values: BigNumberish[];
  fakes: boolean[];
  secrets: BytesLike[];
};

describe("BlindAuction Contract", function() {
  async function deployContract(): Promise<
    [HardhatEthersSigner, BlindAuction]
  > {
    const [deployer] = await ethers.getSigners();
    const factory = await ethers.getContractFactory("BlindAuction");
    const contx = await factory.deploy();
    await contx.waitForDeployment();

    return [deployer, contx];
  }

  function computeHash(bid: BidCommit): string {
    // uniqueness, randomness per called,
    const secret = ethers.toBigInt(ethers.randomBytes(32));
    // more context like [tx id, session id, address]
    const SemiSecret = "semi-secret or fake";

    const encoded = ethers.solidityPacked(
      ["uint256", "bool", "bytes32"],
      [bid.deposit, bid.fake, bid.secret]
    );
    return ethers.keccak256(encoded);
  }

  async function commitBids(auth: HardhatEthersSigner, contx: BlindAuction, bids: BidCommit[]) {
    const authContx = contx.connect(auth);
    for (const bid of bids) {
      const hashCommit = computeHash(bid);
      await authContx.bid(hashCommit, { value: bid.deposit });
    }
  }

  function bidsToRevealArg(bids: BidCommit[]): RevealArgs {
    return bids.reduce(
      (acc, { deposit, fake, secret }) => {
        acc.values.push(deposit);
        acc.fakes.push(fake);
        acc.secrets.push(secret);

        return acc;
      },
      <RevealArgs>{ values: [], fakes: [], secrets: [] }
    );
  }

  it(
    "Should allow commitment, reveal, and correctly determine the highest bidder and refunds",
    async function() {
      const [, contx] = await deployContract();
      const [, oat, art] = await ethers.getSigners();

      // generate off-chain && store
      const oatBids: BidCommit[] = [
        {
          secret: ethers.randomBytes(32),
          fake: true,
          deposit: ethers.parseEther("10"),
        },
        {
          secret: ethers.randomBytes(32),
          fake: false,
          deposit: ethers.parseEther("5"),
        },
      ];
      await commitBids(oat, contx, oatBids)

      const artBids: BidCommit[] = [
        { secret: ethers.randomBytes(32), fake: false, deposit: ethers.parseEther('20') },
        { secret: ethers.randomBytes(32), fake: true, deposit: ethers.parseEther('30') },
      ];
      await commitBids(art, contx, artBids)

      // before reveal, Shouldn't able to see highestBidder, bid.
      {
        const highestBidder = await contx.highestBidder();
        const highestBid = await contx.highestBid();
        assert.equal(highestBidder, ethers.ZeroAddress);
        assert.equal(highestBid, toBigInt(0));
      }

      // after reveal, Should able to
      // - see the highestBidder, bid
      // - revealer receive refund
      // @testcase: oat reveal, should get back refund, art is the highest bidder
      {

        // @reveal
        const revealArgs = bidsToRevealArg(oatBids);
        const provider = ethers.provider;
        {
          const oatWeiBalance = await provider.getBalance(oat.address);
          console.log("before:reveal - eth balance", ethers.formatEther(oatWeiBalance));
        }

        const contxAsOat = contx.connect(oat);
        await expect(contxAsOat.reveal(
          revealArgs.values,
          revealArgs.fakes,
          revealArgs.secrets
        )).to.emit(contxAsOat, "HighestBidIncreased").withArgs(ethers.parseEther("5"));

        {
          const oatWeiBalance = await provider.getBalance(oat.address);
          console.log("after:reveal - eth balance", ethers.formatEther(oatWeiBalance));
        }

        // check highestBidder should be oat
        // Oat shouldn't have any refund
        const highestBidder = await contx.highestBidder()
        assert.equal(oat.address, highestBidder)

        const oatReturnAmt = await contx.pendingReturns(oat.address)
        assert.equal(oatReturnAmt, ethers.parseEther("0"))
      }

      // Art now want to reveal
      {
        const contxAsArt = contx.connect(art)

        const revealArgs = bidsToRevealArg(artBids);
        await expect(contxAsArt.reveal(
          revealArgs.values,
          revealArgs.fakes,
          revealArgs.secrets
        )).to.emit(contxAsArt, "HighestBidIncreased").withArgs(ethers.parseEther("20"));

        const bidder = await contx.highestBidder()
        assert.equal(bidder, art.address);

        const oatReturn = await contxAsArt.pendingReturns(oat.address)
        assert.notEqual(oatReturn, ethers.parseEther("0"));
      }
    });
});
