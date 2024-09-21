# JITLiq Network Node

Liquidity is widespread, often held in spot markets or locked within protocols like Aave. However, locked liquidity can lead to underutilization and become challenging to manage, especially across hundreds of chains, including bridges. The Just-In-Time (JIT) Pool addresses this by removing the need for locked liquidity. Instead, it only provides the required amount precisely when needed, allowing users to retain full access to their assets while still participating in the pool.

The JITLiq Network leverages staking to securely grant access to the pool's liquidity. The process is straightforward: participants stake their assets, which are then synchronized across both the source and destination chains. This synchronization ensures consistency and security throughout the network. Access to the pool is limited to the amount each participant has staked, meaning solvers can only draw liquidity equivalent to their staked amount. This mechanism not only secures the funds but also aligns incentives, ensuring responsible usage of the pool's resources.

![alt text](https://github.com/JITLiq/contracts-core/raw/main/image.png)

[**Technical explainer**](https://bit.ly/JITLiq)

## Protocols Used

-   **LayerZero**: to create an OApp (omnichain app) that communicates across chains for managing bridging & settlement of funds
-   **Circle CCTP**: to bridge funds for reverse settlement of LP funds
-   **Arbitrum, Base, Flow**: contract deployments

## License

[MIT](./LICENSE)
