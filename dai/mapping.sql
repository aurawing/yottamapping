CREATE TABLE `mapping` (
  `transactionHash` varchar(66) NOT NULL,
  `blockNumber` int(10) NOT NULL,
  `ethAddress` varchar(42) NOT NULL,
  `balance` varchar(28) NOT NULL,
  `param` varchar(1000) NOT NULL,
  `status` int(10) NOT NULL DEFAULT '0',
  `isVote` tinyint(10) NOT NULL DEFAULT '0',
  `nodeAccount` varchar(12) DEFAULT '',
  `ytaAccount` varchar(13) DEFAULT '',
  `blockRule` varchar(1000) DEFAULT '',
  `frozenTime` bigint(20) DEFAULT '0',
  `txid1` varchar(200) DEFAULT '',
  `txid2` varchar(200) DEFAULT '',
  `modifyTime` bigint(20) DEFAULT '0',
  PRIMARY KEY (`transactionHash`),
  KEY `BLOCKNUM_INDEX` (`blockNumber` DESC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `test`.`bkrange` (
  `id` INT NOT NULL DEFAULT 1,
  `end` INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
);

INSERT INTO `test`.`bkrange` (`id`, `end`) VALUES ('1', '6800000');