## Governance 컨트랙트

일주일 단위로 안건를 제안하고 결과가 처리되는 컨트랙트 입니다.<br>
제안된 안건에 대해서는 매주 월요일 ( UTC 시간을 기준으로 + 4 days) 투표가 시작되고, 3일동안 투표가 진행됩니다.

## BMERC1155
투표권(VotingPower) 컨트랙트 입니다.<br>
투표권은 매주 갱신해야 합니다. (참여하지 않는 유저의 종적수를 제외하기 위해)<br>
그 주에 해당하는 투표권 ID 는 (현재시간 / 1 weeks) 으로 계산됩니다.

## BMERC20
투표권(ERC1155) 를 얻을 수 있는 기준이 되는 토큰 입니다.<br>
해당 토큰을 소유하고 있다면 그에 해당하는 ERC1155 토큰을 매주 할당 받을수 있습니다.<br>
이 토큰은 최초 구매 할 수 있으며, 투표에 참여하면서 추가적으로 획득할수 있으며, 투표 결과에 따라 차등 지급됩니다.

#### [go-hardhat](https://github.com/bang9ming9/go-hardhat) 을 이용한 프로젝트 입니다.
```bash
# compile & abigen
git submodule update --init # Get openzeppelin-contracts
bms compile --exclude ./contracts/openzeppelin --filter BmErc20,BmErc1155,BmGovernor
```