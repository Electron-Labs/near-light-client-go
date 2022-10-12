package light

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/electron-labs/near-light-client-go/mock"
	"github.com/electron-labs/near-light-client-go/nearprimitive"
	"github.com/near/borsh-go"
)

const (
	LIGHT_CLIENT_BLOCK = `
	{
		"id": "dontcare",
		"jsonrpc": "2.0",
		"result": {
			"approvals_after_next": [
				"ed25519:28DZYLchvppXjTCgxzgJtL2Vbr5gHCJ1Grb1AKSSxNFwiFHdPh6E5bGxcgiFT36KtgvBu8D1CESTVeCxThTMpX5Q",
				"ed25519:iqty9t5qYxwan4WjNYdfti88oC5GMD7xKW12Nayhdz9GFhEfTds15VjRcQEmysGHYH4YaPRoEKdyANkow8XKGs5",
				"ed25519:3SsHV85UuYZwCUUfpgc1bnUQXYEo8Y9Hg8HUVmMbo7aAKvS44iv5s5F7J6msd5236LdZYCF4wt3mqqeQSrnjhMir",
				"ed25519:3bna2miPrNY8yVk9hVeMHUUuETQEumUKmP44AjnJqWfQfoyVHf4oTTTAWSp6SD4A3c1jf9eQ3BLWg7fJN1mpxUtk",
				"ed25519:eKd2G9z3xqJepi9Zkod16GGPDfgTYuf2ppGCVJWKgnUKnMzSQ9HoGzh7jY6aW3ySRmgLhavfqeN2z7FABgi7Sxs",
				null,
				"ed25519:4gQhGGGNz3NscbZ397hrhem1P6KV7Betupd8N5jx8rgYViAsQEAGgz1wwout65uajKRD3RsuSciWUqUByhSNaT6c",
				"ed25519:4ysmVYDSF4tLSKLfYjccJKqthFnLp87nGE7LnmYKmERFwwwk5ZkMcYUVxkPMHZvtLixfcUW59cs8jn4yZB6faq4",
				"ed25519:2ZgXTWKyqbaZ3EV3xBUUyhFraGJufPeD95g4wQPqfcV79nmVkrXqU3d449CAdnbFAGAu5EkspXRT71GXSAh1T1ho",
				"ed25519:5ibY9gznYzpN6zAZjx7S5RqGNsNU3xjyzum14vYxwUiXdntPNjMcLdxq5asFKcZh49rvN3SwLym3NbJSHWmgaRQW",
				"ed25519:3Qk6oLZxD4r6DdjmbpbKQFYVekfVvFYk6HxsyAjfynATor1jJXH7aEy2G1kZNNp7q7QZraePCPTE5vqehp7UJKtP",
				"ed25519:2oWkGtiaa2326Aaj9E3kG2nabgxmwsiLdQVy8kxUpWycAPDEwMVEdhxXzhptwTsE3tmTKQYNWCSVDHxriyANuqjx",
				"ed25519:2jvZNAv7atJEjLiLEXZPFyxKDagwzEpfMEUdnpoGo54sPZ1Mt93qFg8eWPzyT2DgoFwUADTbp5TXjPB2JTypss21",
				"ed25519:2uvqqSibX8jddL6k1CSbCYdmEN8Lr8i1Ji3UuKexVYtUcCrizhNwsXn4CFgVVMCrBFTo1L7yx3SmHViZ4G4XLBH4",
				"ed25519:3LayeVznDXjBLWHMxQ4BPASfMR8EDmf4APxT869r6CPXcCh7p17izAeDT9aA5dmcpJXgLi9GiqHnNLfJ5jKYjnLb",
				"ed25519:1Qeuq1cw7n44Ly1ZqzJn43CfP8S4g9gpaGCzPV5Y38aP7x9wL4RWE4WS6Ha554YWMp8964Y3GgNwjPcP8rxvv7g",
				"ed25519:3bomNpdfcrDtK3t5QepSVxQSrkzLUJUzATTwJRLUqDzZEYXvYvvjMfpuGwQCErYzzJxiuB7dFfxL6p4vXEVgs89x",
				"ed25519:3rma9dot2VoYGUTLX5QMnmkQSxp7msKKZPyvWwXGDdzuGn6eL6yuYjhsmZtctN3eptHbMHcPSkNJPWPAN22u8Uj5",
				"ed25519:2tQ2LBiTm5dK93f35kxnFDFasNWreYd1JHpsBmNc8u61mE6djmbdsTkjPk8qGvoeXjoDk79bWjLsoqKv4H61a3Xr",
				null,
				"ed25519:48FiBja8u93n1YwGZnJV11Gmn7bQjetgn3qXbiQpnvVtC3g9EfXQxywg45pEdhKLry1bZam5sMQWMRMWYnbnkmnX",
				"ed25519:2thqHWuqaYW2Gfo56rFmyd1kzGToHwWDp5JWh3wDzgmY3AKQXSi7uVLEnSYjBuduUyi261JpH2PEC2VW3GgQ4b1s",
				"ed25519:Nw2fG6qUpc847mHjBEot6PnAokogzwNKLZ93vPkuRTc22vUo8gYaRwYhSXe4hi5VcttW8DEVRTUzWzxAAzw5KhB",
				"ed25519:4Bpy8phH7XwfEvFukucFwSC1N73skrafeXzJusfeVNDFJT5NEbbmW5zc17n7RQHiujvucuKjSVsptAuL5tvxp8Zg",
				"ed25519:3rDBs1CXDjqAvME1vxQbMwsPh6c4FvFVC8puY2n8k6SABsPeDQf4cynZm4KwkPiLciqFzKFmi937gejidymFnVaq",
				"ed25519:4gh7NChnFv85nfun95CS9xffSYd9XmLaFtHaEcEyMQrWKn2ynyYTKYLMamLbmUhNwaqJ3c791zW2HupSaJrEtzDb",
				"ed25519:y8XuU7satvG7rMa5C5go3jwy3zh9FpCGzgDmysdqXZn9hmJnUVj8w1gQYDovQgFY8uWCFPDzTjvJzhwAsCQhWnk",
				"ed25519:4QS4L2Mm9WHbGxP5XsUsE5kJJtzz9RZUuY6y8nTYy9KHYpFFCkWN5hb8bJFcHpEi1K12c7p8yy1g73NNZGhRWKCt",
				"ed25519:3XWyULNYTce7qoWEKy4xn5QRpQ5QqJNA4MJwHnsECuHQWhMPYph2MLdgBNiN17D6DgnVh51zMQioM8QX1ugh6NZi",
				"ed25519:5YtCmqoGRM3BqPYNqhZzAyAQqeEXfh9HkjbMpj9iGdEYHV2oauZeryxHvsphgWr3nf92WwBxHjYL2K9b8wRZhMEC",
				"ed25519:5ua7q4FmUBJZksqz2e74S9Uq9B73gQGeGX1bXEkQiEyAobcqzM84cZQ9TXmbxnPp6kus1dtbVfimR7QMjiRchLjf",
				"ed25519:u5pknqrZY37AmRzEsGA77xFe8kAkeLCPVNwTRwSsRYD4vuGX6Wk8VZA5BzM4iN91AWZB4tT1mLR4DEBtB3yYjLs",
				null,
				"ed25519:4Q2yWFGtnJr91cLg5SNif2rEofBmhya4xruktdvCWEXxtG3BTMewTN3pJEPpaPK3TdFw7S3qFbDxG5vJYiD2tcDN",
				"ed25519:3xvaa4CeHyY3eXhhGFnSpj4te6dBPZrSHMsKixpfvL9Uq37YwNTPEZ5QvjvfFrvo1q2MYhSzL82nxQqG5xCTfZL4",
				null,
				"ed25519:5romFpmQYAur711b3oNWfsg3zU9W9NZq5Y5ncTGS8XwsgHakG6NjnjxEzuk2Xsohx5Wy5JfeFPLpkjfvM9pNtKuF",
				null,
				"ed25519:mPj978oVr38YFPWqtfR2bsiGuwduiMCwnV1CkH7QiXCd5LchwNHqoGKr1Lb3CnBUanH8ZRFMHgFVVBCNvdywpnF",
				"ed25519:5jw22tUZGJTNyJWvjtmcBPAmSNVFr8dRvd8nMA1do6RWg5sLREphtnNZTFN2eDRM7XzfWA5xqheCvr8BTtyRhn1u",
				"ed25519:9cA69tCJgoL7EcL8WfeZyU32TyDNkUjewaS2DBFde42fumVvigCzgxutwbr2s4b99B2ssQAZnCkLj888FpfoTVL",
				null,
				null,
				"ed25519:ecsgpnMFfpa6hQGha9J9oWCB6cQPEZg5ryG9N1GrMPhCVFUBhG5w2GhWpnDv8kzPVAuxrJrK17TuHjWLutJPk5q",
				null,
				"ed25519:63KedtRzTF2EvdrkFPuxkdMcm4Gsi3cS8X69Zd36xhWqkbRwXFSz4b544tTQDW7nEjA5eHfa73VatZWpL4nK9ioY",
				"ed25519:4kVfL7yBsBnrP7AA4QxWHY869ACxZSNGAoWnvv6bDBPAb1mAHK9gHB7XfgrQ5oRGFtvMMjywP5rSoabZV2jQbBd",
				"ed25519:2b36Gihq1pUL7WgP6NoWTyKLtjhe3bMLzQV7t2pShXfJtEEC8drtp3V8TaqXLhtKHNehLnh3epfCxyu4Hq2Yx4if",
				"ed25519:61TwgCX1XbwpxDkDkQq1Y9Hd74NjfzYGYj4ezaGrBAhpBjSizEQk3XHuQ81rG4bMiJth9mFeHNpwEjajEXy6M9C1",
				"ed25519:5E6xMNNwzLCSCNpYqg4p58SijQbMoBoAEe2faZXR3wwPUQ16iq5RdouABaRDLMLMYATiYAHged5SSQ7g6WyCEwqL",
				"ed25519:3y3Lme7fqVLxP6jciD3fBk4BP5SXSunwwJwmo7B6KsmvTjXe1znRkPbAg9Sr24nWVY3sm8LWJas1eMX6UrQ9BYso",
				"ed25519:4uPmCptGBUSBziVbwYfPC9GL3KSpLUVaafxxW56wnWhMoEqvug5brA1B4GMpZnYUNsPojjychj517ijCbqVMgCS2",
				"ed25519:3bn5FE9F1bv8eKfU15xDnQt31BC2hV9TmbuhHrusJYMQXbrTFqBzfqYRDotKRuvoxd5NRPhg9GMFJvCSrdzq6NUC",
				null,
				null,
				null,
				null,
				"ed25519:5JcVVEuWYrJMdpbCYXY1hgcMH2toNVW8SRW2V8yqV3bTNAqtyB6UAsyLygnFrorR4u9THZP4mtrJQaoehyTpoGjF",
				null
			],
			"inner_lite": {
				"block_merkle_root": "4xHziXrvtDKjdUjXmGfe5apzeMhMm1CpjJ8dBtRkiiPy",
				"epoch_id": "4Wu9U6C3P9KAAymDYo5W5hv11yi7Xgw6UnyFS6u8V4T9",
				"height": 102422011,
				"next_bp_hash": "hHSgrfXHPYedWpizLYVvhHZ23r1jtEJjYfK4EkqW4mZ",
				"next_epoch_id": "5UkVg2Y7QqDXaBUjDyri7uuq3iiZfmb7FpKpLoET337J",
				"outcome_root": "2DoUX6XsDr5BxRN821ZxTLYYcQBzSSxPMTqMU4TLfu35",
				"prev_state_root": "Fu6AKwXo919gk1Wpti1uBxDzZKrsm96oMfERAkRmLzyF",
				"timestamp": 1665466981632520289,
				"timestamp_nanosec": "1665466981632520289"
			},
			"inner_rest_hash": "9KEuA7Ctb3PTmqvpsjJjQ2wqUGkXLchTdJDxFLiPRMZr",
			"next_block_inner_hash": "3Ce5ULCax6Hqyp3otzvPfm1bryEmHLnRmWBzaza5iHyU",
			"next_bps": [
				{
					"account_id": "node0",
					"public_key": "ed25519:7PGseFbWxvYVgZ89K1uTJKYoKetWs7BJtbyXDzfbAcqX",
					"stake": "31880390716825716686941425088865",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "node3",
					"public_key": "ed25519:ydgzeXHJ5Xyt7M1gXLxqLBW1Ejx6scNV5Nx2pxFM8su",
					"stake": "31880372652970774650650360744540",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "node2",
					"public_key": "ed25519:GkDv7nSMS3xcqA45cpMvFmfV1o4fRF6zYo1JRR6mNqg5",
					"stake": "31880308326225864071188332802669",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "node1",
					"public_key": "ed25519:6DSjZ8mvsRZDvFqFxo8tCKePG96omXW7eVYVSySmDk8e",
					"stake": "31878084085605515853527117453930",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "cryptogarik.pool.f863973.m0",
					"public_key": "ed25519:FyFYc2MVwgitVf4NDLawxVoiwUZ1gYsxGesGPvaZcv6j",
					"stake": "14305979176208732721121145425795",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "aurora.pool.f863973.m0",
					"public_key": "ed25519:9c7mczZpNzJz98V1sDeGybfD4gMybP4JKHotH8RrrHTm",
					"stake": "13300893885949214473679563546988",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "01node.pool.f863973.m0",
					"public_key": "ed25519:3iNqnvBgxJPXCxu6hNdvJso1PEAc1miAD35KQMBCA3aL",
					"stake": "9439666705123769983372375372648",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "legends.pool.f863973.m0",
					"public_key": "ed25519:AhQ6sUifJYgjqarXSAzdDZU9ZixpUesP9JEH1Vr7NbaF",
					"stake": "7855048372135223531280830598430",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "spectrum.pool.f863973.m0",
					"public_key": "ed25519:ASecMN9e28vtCJn7rD2noNwL5c3odzQgAfbfHrUnbSVe",
					"stake": "7356598107221425952833976770890",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "nodeasy.pool.f863973.m0",
					"public_key": "ed25519:25Dhg8NBvQhsVTuugav3t1To1X1zKiomDmnh8yN9hHMb",
					"stake": "7194100699895647377582537678219",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "everstake.pool.f863973.m0",
					"public_key": "ed25519:4LDN8tZUTRRc4siGmYCPA67tRyxStACDchdGDZYKdFsw",
					"stake": "7144066381969590047312613793768",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "chorusone.pool.f863973.m0",
					"public_key": "ed25519:3TkUuDpzrq75KtJhkuLfNNJBPHR5QEWpDxrter3znwto",
					"stake": "6891410601843231354846544118820",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "ni.pool.f863973.m0",
					"public_key": "ed25519:GfCfFkLk2twbAWdsS3tr7C2eaiHN3znSfbshS5e8NqBS",
					"stake": "6640497327496677469514313949325",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "stakely_v2.pool.f863973.m0",
					"public_key": "ed25519:7BanKZKGvFjK5Yy83gfJ71vPhqRwsDDyVHrV2FMJCUWr",
					"stake": "6235550137617161986397344541506",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "staked.pool.f863973.m0",
					"public_key": "ed25519:D2afKYVaKQ1LGiWbMAZRfkKLgqimTR74wvtESvjx5Ft2",
					"stake": "4675501969510072713584269203958",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "foundryusa.pool.f863973.m0",
					"public_key": "ed25519:ABGnMW8c87ZKWxvZLLWgvrNe72HN7UoSf4cTBxCHbEE5",
					"stake": "1703002133636428149117892943816",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "tribe-pool.pool.f863973.m0",
					"public_key": "ed25519:CRS4HTSAeiP8FKD3c3ZrCL5pC92Mu1LQaWj22keThwFY",
					"stake": "1642282732964257844219728804768",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "sweden.pool.f863973.m0",
					"public_key": "ed25519:2RVUnsMEZhGCj1A3vLZBGjj3i9SQ2L46Z1Z41aEgBzXg",
					"stake": "1566741432877745197229705088539",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "chorus-one.pool.f863973.m0",
					"public_key": "ed25519:6LFwyEEsqhuDxorWfsKcPPs324zLWTaoqk4o6RDXN7Qc",
					"stake": "1537228829801614061234008187779",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "hotones.pool.f863973.m0",
					"public_key": "ed25519:2fc5xtbafKiLtxHskoPL2x7BpijxSZcwcAjzXceaxxWt",
					"stake": "1315292167927479480381218976652",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "pathrocknetwork.pool.f863973.m0",
					"public_key": "ed25519:CGzLGZEMb84nRSRZ7Au1ETAoQyN7SQXQi55fYafXq736",
					"stake": "1052411613249714898795964424000",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "stakesstone.pool.f863973.m0",
					"public_key": "ed25519:3aAdsKUuzZbjW9hHnmLWFRKwXjmcxsnLNLfNL4gP1wJ8",
					"stake": "906333851427798828850956526201",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "leadnode.pool.f863973.m0",
					"public_key": "ed25519:CdP6CBFETfWYzrEedmpeqkR6rsJNeT22oUFn2mEDGk5i",
					"stake": "902683060684666906409010959459",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "dsrvlabs.pool.f863973.m0",
					"public_key": "ed25519:61ei2efmmLkeDR1CG6JDEC2U3oZCUuC2K1X16Vmxrud9",
					"stake": "898818019655781395004122901699",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "blockscope.pool.f863973.m0",
					"public_key": "ed25519:6K6xRp88BCQX5pcyrfkXDU371awMAmdXQY4gsxgjKmZz",
					"stake": "893255585471702649570152737975",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "al3c5.pool.f863973.m0",
					"public_key": "ed25519:BoYixTjyBePQ1VYP3s29rZfjtz1FLQ9og4FWZB5UgWCZ",
					"stake": "892492981005236348609258542340",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "grassets.pool.f863973.m0",
					"public_key": "ed25519:3S4967Dt1VeeKrwBdTTR5tFEUFSwh17hEFLATRmtUNYV",
					"stake": "890477852030951358598788920485",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "shurik.pool.f863973.m0",
					"public_key": "ed25519:9zEn7DVpvQDxWdj5jSgrqJzqsLo8T9Wv37t83NXBiWi6",
					"stake": "856795480920515766403790972105",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "baziliknear.pool.f863973.m0",
					"public_key": "ed25519:9Rbzfkhkk6RSa1HoPnJXS4q2nn1DwYeB4HMfJBB4WQpU",
					"stake": "853137540071718175183075355279",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "optimusvalidatornetwork.pool.f863973.m0",
					"public_key": "ed25519:BGoxGmpvN7HdUSREQXfjH6kw5G6ph7NBXVfBVfUSH85V",
					"stake": "812700276671099460216641999283",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "chelovek_iz_naroda.pool.f863973.m0",
					"public_key": "ed25519:89aWsXXytjAZxyefXuGN73efnM9ugKTjPEGV4hDco8AZ",
					"stake": "805435787718810793997158732391",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "basilisk-stake.pool.f863973.m0",
					"public_key": "ed25519:CFo8vxoEUZoxbs87mGtG8qWUvSBHB91Vc6qWsaEXQ5cY",
					"stake": "782212199058438713768164602174",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "ou812.pool.f863973.m0",
					"public_key": "ed25519:2APjYBPnQ7CGDxFpsUHeAcyhpjZRWWpjciq2Pdzk31uQ",
					"stake": "631366416135881486533604335885",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "dimasik.pool.f863973.m0",
					"public_key": "ed25519:3gqyuPas4axMAMz4VEKF7cSxT9ZGnJfpLbzGZZ61mZvU",
					"stake": "501862499014281739520614986286",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "blackhox.pool.f863973.m0",
					"public_key": "ed25519:3jqcMLsco4aMLtWr35KMEw5W5z4G9TkcYvR5btfHipn9",
					"stake": "477144965033117573566249889543",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "blazenet.pool.f863973.m0",
					"public_key": "ed25519:DiogP36wBXKFpFeqirrxN8G2Mq9vnakgBvgnHdL9CcN3",
					"stake": "473288656246900936904440658707",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "infstones.pool.f863973.m0",
					"public_key": "ed25519:BLP6HB8tcwYRTxswQ2YRaJ5sGj1dgGpUUfcNwbnWFGCU",
					"stake": "471885822894344396706821471143",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "stingray.pool.f863973.m0",
					"public_key": "ed25519:9sTjViLyTuaBe8LEX341aB8iRd6tGdpKgiv6jEiUxPgQ",
					"stake": "380317603658395718684208271366",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "kiln.pool.f863973.m0",
					"public_key": "ed25519:Bq8fe1eUgDRexX2CYDMhMMQBiN13j8vTAVFyTNhEfh1W",
					"stake": "256360714051193755808487919792",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "idtcn3.pool.f863973.m0",
					"public_key": "ed25519:DtkY9WtkWweSrF13BJi5k4c6xyk3tBAC9y92AEY4Ayfb",
					"stake": "154909524382635392035218055051",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "shardlabs.pool.f863973.m0",
					"public_key": "ed25519:DxmhGQZ6oqdxw7qGBvzLuBzE6XQjEh67hk5tt66vhLqL",
					"stake": "88716568211908202592232235256",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "gettingnear.pool.f863973.m0",
					"public_key": "ed25519:5QzHuNZ4stznMwf3xbDfYGUbjVt8w48q8hinDRmVx41z",
					"stake": "77361921161924967114893017917",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "leadnode-shard.pool.f863973.m0",
					"public_key": "ed25519:CzWox9TE1AR4xfDracD5eN4xy5f91ZNu5CTRsPcdH45C",
					"stake": "73241234958550983459388379881",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "sevennines-t0.pool.f863973.m0",
					"public_key": "ed25519:BHKMMc1t7F6B26BbaBVMhex3riPtWoGL5CricgXiSFC4",
					"stake": "63516206635776140113080771331",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "redhead.pool.f863973.m0",
					"public_key": "ed25519:5qPvLhc86TDdof4YEjBMKrENzT3UA9mEonKWQXuFvaHX",
					"stake": "56393870147050517009727307224",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "azetsi.pool.f863973.m0",
					"public_key": "ed25519:2MFKLj9E2kRdJoQqUgaY9KtheebzLv9ntdgTGsZxLaE1",
					"stake": "50377534576586039354769572795",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "wackazong.pool.f863973.m0",
					"public_key": "ed25519:EK1bdY5F6prLush2aKnJBe5neHdK52wwTjg3qY4v3cjX",
					"stake": "47751441525930200072873587519",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "sergo.pool.f863973.m0",
					"public_key": "ed25519:3uV2DGyNVfSgEPi9UdwFWBHtLZyPHmsN4iNtF7nEvZnT",
					"stake": "47696413993663018631748398653",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "guardia.pool.f863973.m0",
					"public_key": "ed25519:2b5AQqcf8PHAUzqxWYoaFiTEa7QuEkpihzeoDittQaPL",
					"stake": "46604814168002295580519267172",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "tandemk.pool.f863973.m0",
					"public_key": "ed25519:8zqx8dzqsxXivPMcSk2e7gezMvooZb1RWMGHA1FzFrak",
					"stake": "46572247940000099754391496222",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "pero_val.pool.f863973.m0",
					"public_key": "ed25519:J43JCHe2XKU7wiDi7PSS8exsdVgnYHcbnC8R14SpF6HV",
					"stake": "45761901930241420798583789391",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "pandateam.pool.f863973.m0",
					"public_key": "ed25519:7n426KJocZpJ5496UHp6onwYqWyt5xuiAyzvTGwCQLTN",
					"stake": "45385922119728145519829756542",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "meduza.pool.f863973.m0",
					"public_key": "ed25519:HPNwQG4gabi389RQfyj7ZN2o7Y8BjsU3gnyataubHZTk",
					"stake": "44874623228874793616912655096",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "adel0515.pool.f863973.m0",
					"public_key": "ed25519:Bo7FzbbrRGRWagU58Fx1i8zFjN154rSA5nFkt6urHLYh",
					"stake": "44810969235385841571995198233",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "gruberx.pool.f863973.m0",
					"public_key": "ed25519:9HpDAKMwmUBuME7Zmg77rrdQVixEKeKPni8HAJBCAxwd",
					"stake": "44095045919300306477073947468",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "nodebull.pool.f863973.m0",
					"public_key": "ed25519:8tgk15x5XU15Ka9UCVCmAxeJweRrTbA96rnigs2ARQP4",
					"stake": "41723876554385454899808377772",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "p2pstaking.pool.f863973.m0",
					"public_key": "ed25519:A13kqREA5eu35SXyB6TH2j6HjVyyCUEJr8yvkikKV9Wt",
					"stake": "41610132302436462827176012554",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "makil.pool.f863973.m0",
					"public_key": "ed25519:CCuzQ9CE4HLXMqE55e5jKNdfjVSrskQnNe3LPoSeeHNj",
					"stake": "40494972222691784686464364864",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "kt2.staking-farm-factory.testnet",
					"public_key": "ed25519:HzJtjiAzXrhG6mGGgNCLx4JTivaP6hwssd7gD2ve2Ej6",
					"stake": "40030000830373239417600000000",
					"validator_stake_struct_version": "V1"
				},
				{
					"account_id": "kuutamo.pool.f863973.m0",
					"public_key": "ed25519:8T7J4vNjoUkQ8auYiqkxofELaa18aMGc4qNhCEx7qHCg",
					"stake": "39860917805431675582724281535",
					"validator_stake_struct_version": "V1"
				}
			],
			"prev_block_hash": "8A3UaDA8ZjGGvVg5tDtGAkf46LrXQkamEsSoHG8LQ3v5"
		}
	}`

	EXECUTION_OUTCOME = `
	{
		"id": "dontcare",
		"jsonrpc": "2.0",
		"result": {
		  "block_header_lite": {
			"inner_lite": {
			  "block_merkle_root": "CRVMDaFCLz5GDKtgRzEqi2Rde52yzEDbLabtC2jK7nZm",
			  "epoch_id": "EeC8QHiPSdr6CSDhJiCQL4wMR8or33qvkvirzh9Moe6x",
			  "height": 102367480,
			  "next_bp_hash": "74P742gjuiU6UTxpzgPR1L4c1iqMu6ZtPxFj656XAyCx",
			  "next_epoch_id": "4Wu9U6C3P9KAAymDYo5W5hv11yi7Xgw6UnyFS6u8V4T9",
			  "outcome_root": "3yq51ESCg5st9qk7aksomjFc3hQoL2dobUdKg6TmshT9",
			  "prev_state_root": "EB8aWEHdXVomTwJZFgsTkRsVCk31fw2aqSxkL6R5eu6b",
			  "timestamp": 1665413019091964617,
			  "timestamp_nanosec": "1665413019091964617"
			},
			"inner_rest_hash": "FrHB6FJo8c8cPt3fVGz7QdfKfdwSXB3QWHkfpAeMDRzF",
			"prev_block_hash": "YUK3BcpAx3MvXtfgsqUTeRCq3tNnZW2N6xL3zcpJAAM"
		  },
		  "block_proof": [
			{
			  "direction": "Right",
			  "hash": "HqLoC4DL4mKZWoAuThT4tBKoS5qwPUmYx9UYZZMRbvms"
			},
			{
			  "direction": "Right",
			  "hash": "E4HJmteNwLvVzLAyo1C88xj4vb3TkkBFCHfEXgCkNXhN"
			},
			{
			  "direction": "Right",
			  "hash": "CwafrqqpdnQPUEUGCTCeZLZGbb2Xa1oUWKg9x7T71w43"
			},
			{
			  "direction": "Left",
			  "hash": "HJQCSyvJFdMF3Ua9tX9HTTBhf3ZXu8tXT98ZD8YRysVs"
			},
			{
			  "direction": "Right",
			  "hash": "DAMyXn1Gp1WDpQLDy9bxK1M3Euyjqb3kTSmwL42gnnYG"
			},
			{
			  "direction": "Right",
			  "hash": "8ad4uvqVVZya8D8rEfuJSt7UW6A9VhmZ3QGNm6P3zL1Z"
			},
			{
			  "direction": "Left",
			  "hash": "8jj7nPKuhkycSJXodX2Ajqi7XkPpPsStPjZpJGZtM2Uq"
			},
			{
			  "direction": "Right",
			  "hash": "87etPUFfUt2ybiMu2HbLBomMNVvdccZbpAQJoEhyEdXf"
			},
			{
			  "direction": "Left",
			  "hash": "2SRQ6q9ZtXf9EdFU95TCaCSsNDppAMtCMo1ZJnwfzP3i"
			},
			{
			  "direction": "Left",
			  "hash": "9aQtyk3aX1gqhBcv9YB1z2QsqmNuGKSWUj5SQUG8HvNg"
			},
			{
			  "direction": "Left",
			  "hash": "9qKmfMeLGWwxyFGGu1vyVdEGMKb3St185stmLq8rzgUh"
			},
			{
			  "direction": "Right",
			  "hash": "8PRLnyMWLFbK1toSGYQRZpidKY78h4d4rmWQ3ArJv7NH"
			},
			{
			  "direction": "Right",
			  "hash": "2jTjMrA4Cm6pRD8taKsu7K1R4xbLyqyWbYZfVMSYpeJW"
			},
			{
			  "direction": "Right",
			  "hash": "7w2UGWw5wveZTJmVW2MHoy3rMUKMhaP4GrGWz1AC1Bc8"
			},
			{
			  "direction": "Left",
			  "hash": "9cJ8taXvoAedWXZUFiJ7ZYijdhoTeUSYoK3oUFqJLUat"
			},
			{
			  "direction": "Left",
			  "hash": "EEs3nofauM8tN7zWoVjvXpK2EVqKMVrbGqGpyjL1DAKn"
			},
			{
			  "direction": "Right",
			  "hash": "5PUSvGkte8QtBvvyGz9mwceXtB88fjjTZcZQ8jQAP438"
			},
			{
			  "direction": "Left",
			  "hash": "3eask1LZa87NuXFYUCAJuWWvgLvj2rpt5ybM3VxC5vne"
			},
			{
			  "direction": "Left",
			  "hash": "4xiVKFyAaKx2JSNGZ3xAojPssGwoVwNSsroFFZHjKvak"
			},
			{
			  "direction": "Left",
			  "hash": "yxSr5HD8zXpYFVeWbE6bm88wYautgoFL3qysegeVDHH"
			},
			{
			  "direction": "Left",
			  "hash": "Bkcnkncp3gtWiciCK8QV2pk57MP6wWyBaRw2E7s7WbVb"
			},
			{
			  "direction": "Left",
			  "hash": "76U6DMh4J4VB5sfVVNRpSTeB4SEVt4HPqhtQi2izGZxt"
			}
		  ],
		  "outcome_proof": {
			"block_hash": "EdnpBxt2QyAHvseR8whrbCZXipor8VaQckLiJStUZduv",
			"id": "89aUfq2SU6ktdjtvU6kTCtsueQomgZ7s3dCdoHDZgfrd",
			"outcome": {
			  "executor_id": "partht.testnet",
			  "gas_burnt": 2428117762192,
			  "logs": [],
			  "metadata": {
				"gas_profile": null,
				"version": 1
			  },
			  "receipt_ids": [
				"AYDomG3TpmssBCp2F15qcos9CtharnDHHefKdxToH8Uf"
			  ],
			  "status": {
				"SuccessReceiptId": "AYDomG3TpmssBCp2F15qcos9CtharnDHHefKdxToH8Uf"
			  },
			  "tokens_burnt": "242811776219200000000"
			},
			"proof": [
			  {
				"direction": "Right",
				"hash": "9s8qU4s1aDdBHcby9YZegdh7kGbZb3XCZsGTyAT8RcaA"
			  },
			  {
				"direction": "Right",
				"hash": "AxXVTyGrvZt3YXG9GGZCGDZSejP11Bjt5kJSf1XQ4HFD"
			  }
			]
		  },
		  "outcome_root_proof": [
			{
			  "direction": "Left",
			  "hash": "2oKiuAQnqLhuZ2EbokXL6vsriE7tVCZkYkD1ee2EeDsy"
			},
			{
			  "direction": "Left",
			  "hash": "4A9zZ1umpi36rXiuaKYJZgAjhUH9WoTrnSBXtA3wMdV2"
			}
		  ]
		}
	  }
	`
)

func TestMerkleRootEquality(t *testing.T) {
	nlc_json := NearLightClientBlockView{}
	err := json.Unmarshal([]byte(LIGHT_CLIENT_BLOCK), &nlc_json)
	if err != nil {
		t.Errorf("Failed to parse light client block: %s", err)
	}

	tx_proof_json, err := GetTxProof(EXECUTION_OUTCOME)
	if err != nil {
		t.Errorf("Failed to parse tx proof: %s", err)
	}

	tx_proof, err := tx_proof_json.Parse()
	if err != nil {
		t.Errorf("Failed to parse tx_proof: %s", err)
	}

	ser_il, err := borsh.Serialize(tx_proof.BlockHeaderLite.InnerLite.ToBlockHeaderInnerLiteViewFinal())
	if err != nil {
		t.Errorf("Failed to serialize: %s", err)
	}

	h := mock.MockHostFunction{}

	il_sha := h.Sha256(ser_il)
	re := CurrentBlockHash(h, il_sha, tx_proof.BlockHeaderLite.InnerRestHash, tx_proof.BlockHeaderLite.PrevBlockHash)

	nlc := nlc_json.Parse()

	root, err := Compute_root_from_path(h, tx_proof.BlockProof, nearprimitive.MerkleHash(re))
	if err != nil {
		t.Errorf("Failed to compute root: %s", err)
	}

	if !bytes.Equal(nlc.InnerLite.BlockMerkleRoot[:], root[:]) {
		t.Errorf("Failed to validate tx!")
	}
}
