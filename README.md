# acorus


<!--
parent:
  order: false
-->

<div align="center">
  <h1> acorus repo </h1>
</div>

<div align="center">
  <a href="https://github.com/bridge-alchemy/acorus//releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/bridge-alchemy/acorus.svg" />
  </a>
  <a href="https://github.com/bridge-alchemy/acorus//blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/bridge-alchemy/acorus.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/bridge-alchemy/acorus">
    <img alt="GoDoc" src="https://godoc.org/github.com/bridge-alchemy/acorus?status.svg" />
  </a>
</div>

Acorus is project which can sync l1->l2 deposit, l2->l1 withdraw transactions of all layer2. The back-end service, written in golang, provides http api for upper-layer service docking.

**Tips**: 
- need [Go 1.21+](https://golang.org/dl/)
- need [Postgresql](https://www.postgresql.org/)


## Support layer2 as follow
- ✅ [Optimism](https://github.com/ethereum-optimism)
- ✅ [Scroll](https://github.com/scroll-tech)
- ✅ [Linea](https://github.com/Consensys)
- ✅ [Base](https://github.com/base-org)
- ✅ [Mantle](https://github.com/mantlenetworkio)
- ✅ [Metis Andromeda](https://github.com/MetisProtocol)
- ✅ [Manta](https://pacific.manta.network/)
- ✅ [Polygon](https://github.com/0xPolygonHermez)
- ✅ [Zksync](https://github.com/matter-labs)
- ✅ [Arbitrum](https://github.com/OffchainLab)
- ✅ [ZKFair](https://zkfair.io/)
- 🏗️ [Starknet](https://github.com/starkware-libs)

If you want support your project, please create pr for us, we will support it.


## Install

### Install dependencies
```bash
go mod tidy
```
### build
```bash
cd acorus
make
```

### Config env

- For layer2 chain contracts config, you can [config](https://github.com/bridge-alchemy/acorus/tree/main/config) director and refer to exist config do it.
- For toml config, you can [acorus.toml](https://github.com/bridge-alchemy/acorus/blob/main/acorus.toml) file and config your real env value.

### start syncer
```bash
./acorus syncer
```

### start api
```bash
./acorus api
```

## Contribute

TBD
