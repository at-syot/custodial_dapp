import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@solarity/hardhat-gobind"
import '@typechain/hardhat'
import '@nomicfoundation/hardhat-ethers'
import '@nomicfoundation/hardhat-chai-matchers'

const config: HardhatUserConfig | any = {
  solidity: "0.8.28",
  typechain: {
    outDir: "./typechain-types",
    target: "ethers-v6",
    alwaysGenerateOverloads: false,
  },
  gobind: {
    outdir: "./generated-types/bindings",
    deployable: true,
    runOnCompile: false,
    abigenVersion: "v1",
    verbose: false,
    onlyFiles: [],
    skipFiles: [],
  },
};

export default config;
