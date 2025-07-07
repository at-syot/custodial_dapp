import { expect } from "chai";

describe("todo contract", function() {
  it("deployment ok", async function() {
    const [owner] = await ethers.getSigners();
    const token = await ethers.deployContract("Todo");
    await token.add("hello world");
    await token.add("task 0");
    const what = await token.add("task 1");
    const tasks = await token.list();

    // console.log(what);
    expect(tasks.length).to.equal(3);
  });
});
