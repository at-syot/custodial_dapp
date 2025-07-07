import { expect, assert } from 'chai';
import { ethers } from 'hardhat';
import { toBigInt, type EthersError } from 'ethers'
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers"
import type { Ballot } from '../typechain-types/index'

describe("Ballot Contract", function() {
  async function deployContract(proposals: string[]): Promise<[HardhatEthersSigner, Ballot]> {
    const [owner] = await ethers.getSigners();
    const factory = await ethers.getContractFactory("Ballot")
    const contx = await factory.deploy(proposals)
    await contx.waitForDeployment();

    return [owner, contx];
  }

  it("only chairman can give voter's right", async function() {
    const proposals: string[] = [
      ethers.encodeBytes32String("proposal 1"),
      ethers.encodeBytes32String("proposal 2"),
    ];
    const [, contx] = await deployContract(proposals);
    const [, addr0, addr1] = await ethers.getSigners();

    const txWithChairman = contx.giveRightToVote(addr0.address)
    await expect(txWithChairman).not.to.be.reverted

    const contxOnAddr0 = contx.connect(addr0);
    const txWithAddr0 = contxOnAddr0.giveRightToVote(addr1.address);
    await expect(txWithAddr0).to.be.reverted
  });

  it("should have some proposal after deployment", async function() {
    const proposals: string[] = [
      ethers.encodeBytes32String("proposal 1"),
      ethers.encodeBytes32String("proposal 2"),
    ];
    const [deployerAddr, contx] = await deployContract(proposals);

    try {
      const chairperson = await contx.chairperson();
      assert.equal(chairperson, deployerAddr.address);

      const p = await contx.proposals(0);
      assert.isNotNull(p);
    } catch (e) {
      const ethersErr = <EthersError>e;
      console.error("catch err code - ", ethersErr.code)
      console.error("catch err name - ", ethersErr.name);
      console.error("catch err innner e - ", ethersErr.error);
      console.error("catch err info - ", ethersErr.info);
      console.error("catch err msg - ", ethersErr.message);
      console.error("catch err smsg - ", ethersErr.shortMessage);
      console.error("catch err stack - ", ethersErr.stack);
      // if (ethers.isError(e, "REPLACEMENT_UNDERPRICED")) { }
    }
  })

  it("Should correctly transfer voting weight after a voter delegates their vote",
    async function() {
      const [_, contx] = await deployContract([]);
      const [, oat, art, yok] = await ethers.getSigners();

      await contx.giveRightToVote(oat.address)
      await contx.giveRightToVote(art.address)
      await contx.giveRightToVote(yok.address)
      expect(await contx.voters(oat.address)).to.not.null;

      const contxAsOat = contx.connect(oat);
      const contxAsArt = contx.connect(art);
      await contxAsOat.delegate(art.address);

      // after Oat delegate vote -> Art
      const oatVoter = await contx.voters(oat.address);
      assert.equal(oatVoter.voted, true)
      assert.equal(oatVoter.delegate, art.address)

      const artVoter = await contxAsArt.voters(art.address)
      assert.equal(artVoter.weight, toBigInt(2));
      assert.equal(artVoter.voted, false);
    });

  it("Should correctly determine the winning proposal after multiple votes",
    async function() {
      const proposals: string[] = [
        ethers.encodeBytes32String("proposal 1"),
        ethers.encodeBytes32String("proposal 2"),
      ]
      const [, contx] = await deployContract(proposals);
      const [, oat, art, yok] = await ethers.getSigners()
      await contx.giveRightToVote(oat.address)
      await contx.giveRightToVote(art.address)
      await contx.giveRightToVote(yok.address)

      const contxAsOat = contx.connect(oat)
      const contxAsArt = contx.connect(art)
      const contxAsYok = contx.connect(yok)
      contxAsOat.vote(0)
      contxAsYok.vote(0)
      contxAsArt.vote(1)

      const winningName = await contx.winningName()
      assert.equal(winningName, ethers.encodeBytes32String("proposal 1"))
    })

  it("Should revert tx, when face delegation loop", async function() {
    const [, contx] = await deployContract([]);

    // make delegation loop
    // make tx revert

    const [, oat, art, yok] = await ethers.getSigners()
    await contx.giveRightToVote(oat.address)
    await contx.giveRightToVote(art.address)
    await contx.giveRightToVote(yok.address)

    const contxAsYok = contx.connect(yok)
    const contxAsArt = contx.connect(art)
    const contxAsOat = contx.connect(oat)

    await contxAsYok.delegate(oat.address)
    await contxAsArt.delegate(yok.address)
    await expect(contxAsOat.delegate(art.address))
      .to
      .be
      .revertedWith("We found a delegation loop")
  });
})
