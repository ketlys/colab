package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"flag"
	"log"
	"math"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"
	  "fmt"
	  "os"

	"github.com/valyala/fasthttp"
)

var (
	processes = new(sync.Map) // here we'll store routines statuses

	// everything here is pretty standard, if something doesn't work you should check your network first
	client = &fasthttp.Client{
		MaxConnDuration: time.Second * 30,
		ReadTimeout:     time.Second * 30,
		WriteTimeout:    time.Second * 30,
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, time.Second*5)
		},
	}

  	mu sync.Mutex
  	f *os.File
	
	addresses = []string{
		"DdbHBaGVZqjabJvdnyLyEgimcV2UyUBP24V7aX5cJqaje",
		"DvR3a91B1TcBDjJuWbJEctYqZtgmthJQLDoXijFYfqMYu",
		"Da5trdwuti3Gsz9UZUMd9LkP641VfEFDXTgz7MQAwkxhZ",
		"DzE6wZgSmkhYUx2iBunBoBE6J1bdAdioVSt5ARcLhTZeD",
		"DguzQSkjS8VvRjnj1oB2MtKR7RhRfw45e278e22TdcnsU",
		"E2uLqU2Bp9XKqepsAe9Uq3tUTtSoSQK7eD7pXER9yhidb",
		"Ddtw7d78pPE3wVRYvH9bHZTH8pDwfHBJ8usJQr3wG5JgA",
		"E34hYfyWpbN8hN9iwqiwNVMbv4XwNmqEUjJFTsBosFyPu",
		"DkeGKJKf5JdomprNxJyaCrSaZL5u8Un2n1h5BRYeiwHuv",
		"DzCrNR4p2QeVF5Sjp757PCbQZYrXPoz8m2dfeara9d9Wz",
		"DpyCABs1gaLgaqGWJFQ83U8fkniRN12xFkJ3Y5eyAMhs3",
		"De8Zp41jvVtVodrGFg3LJ7F4cyKkFpQEfUtzj1Ha5U85d",
		"DaDwfDRmByEbT3GjDnHZNdQLV7ibhBh45QVt1xcxHtcM6",
		"DZ8ZtudN6tM3VbXJBD4XpzCdoTnvjnRfmxjC7jD1fbNNn",
		"DfbQ73XSYGj9PACBMGn4kA96qVppWXFjwaPvncFTZkVYj",
		"Dfm17kVoFWgKP4Pg9puvZF8e9B5znYZn1mxYVcpPPyT5N",
		"DYXG5hYMUxEqX8YHoBcMKrEjMKS6Kf5RPhqyx1es12vqz",
		"DwbeUV1MudoMdXx8jpea9jiTVWYjJ3jcKZRrYfErxhLrN",
		"DVZWQZfw5op5Wywm6pE3X7UY67yJDiYqCr4Kkp4NbpjuH",
		"DffhsFiRzXeXcqUQhrXjE9Kmc5CZwFtoRXDGd3DshkUp8",
		"DX9fu1s6xq7uSkn1GJuLXk8h2Hi2V54agR25K5YNmyxtD",
		"DYvXuj1cgQpM9wdwoh5Gcrh9cuynjHKEKDB6fnVHUDZeR",
		"DsrQVK4TLYs8i75aPcixZfVHDNWWE5R2SmNH9bwZtwz6t",
		"DgGNzT9ieWB6Xq9yhXs2MThqs4Y2Qm5a66JCRE7UqAuWp",
		"Dn5uZrPRfGHEc9v4fSdy7ULc5AxLUrpTvVjcNezUm3zQC",
		"DgNPER79YPLxqgqJaPM3yEa6n9iWBRCbEkVrEGSm3m7ft",
		"DaPcvEFuVyFHqKNPGsXZdFCme9Q63PGaHRC2U2tPWBP3x",
		"Dm2ztZtjLiYPSU9hWpUuzH8iMabQckDK97ahr5JnXgzxA",
		"DtqBDfFLKYLnUTL1dEU2nWusQiR1kLPzrJYxZvd5JvEHH",
		"DmGyCZ4xkgxrrGsPTyDzSk9xdJvTNCTxkt9FuyLEqk2NA",
		"E2qAH9XPp3PwA1RXG5zXVHMWw1DuiAKdnudQRGeiPPahK",
		"DzXoghB2AbyVgX7uAiS1bXmnhtCHpqy3gw3AaG6yCkaT4",
		"DswmbNKUMjSUYxzUxAZwejFYdHERGchREbUXCLTm4avrT",
		"DVmGfFuR7vTcdpiDWyQdrdzXitV92YEixpETBoiMVJ1cA",
		"DhCFtnMZdcUr77rzpGXct5zpHSF75h6WWy2ZFMK8hxU2M",
		"DmvM1Z75HTL6hZJ8w7KNNWPRn43eMryznKX5icCD4kAaj",
		"E2Q6eHyfc7k1TA3wn2sXC7EWap9tZgzFXQoNxg7Q8A5PQ",
		"DhCZinJakU64CSDSKMzpozikK6j5Daxzz7ptJxUX66Cqu",
		"E1MMvkKXSERfQ1skEAqCrNLrQ3PWL7CJ3sKftF1qsbrcj",
		"DX5nYwCB1XNoD1QzbNXCPM647Vtxwd6exUacnrmtPdTch",
		"E2XZ6JetMreNVEdZJ4e1J34BL7ATGwyhHUssG2mQjsDrM",
		"DzPgV1u9k7QpCsjK2wh9eoMjJwd6Z9aFNw8PGHfVd4rYT",
		"DqXKeP8H85CtNvMKWHR5NKPXdcsJtranqwmRY81TFDkXk",
		"DuZAza8kUQ5VURo2Ynb7QkUJeeV8gFvqvHazcYfy4kv2b",
		"DgTLdpQk3HfAYYo7Qh2xptbm8KWG2Xy24uxsSijyn4LGn",
		"Ddn5FATg7vSeBeUrSdotzbZrNVidE52pH6B8bUyRfG9cy",
		"DxcMrCPw3aZAZ5s3kXt3Dbu6jXQ7SDx9Ux35YPAigBqvE",
		"DbqjjYSoeUHtn1NFfyPQ4L289ee1sN94V8uYRDWPiBHJF",
		"DwdbfhHF1kFVY1r9Cd8vrVUVtHfCjytXwNAjXc9Xx6y1U",
		"DroGon1b233zxg9oEW6BJQNF6QyWkHaatpk4DDJbYFtTh",
		"E1vDJMXyg1QvGWNFGFQBaiZnGDCofZ9uq6GEF3sb9zFGr",
		"DW9EyZSh7TfMGwkt88uMms7m4v4sjd7sYiSJBEZW38WDt",
		"Dasf9rRV22Ua4XPZdGT7vCXXpE2Vs18bNnAffP47khwYy",
		"E3tWiGzo7TpGFRRy5vJiZGBfdNCoefwcSCoRFhuNgXdAa",
		"DcSNfiRJQ1mj9b36mjX8WXcv1KN54NJf3CSQfnubbAxEg",
		"DdqG4NvNYdh7zZQdpPbL6pofSxwVAzoYVKX9htX8GEE2M",
		"DWcCY7iSMusSV98q5FP3dsnPiTphyQCcr5T2he2iMDJ34",
		"DgKnA5z1NrK58yktxY6kurSuCMMB64L2bFBTxirTAX236",
		"Dn8kt27oC46QuBcZezFMyVvvPWdA5S7QyskgywiQ7s36m",
		"DwygWTiSTZu2fCpZHZPTXuMRtaF1Q8mUSRcYrcP5d2utd",
		"DjWb9TB9NnQ95NvLMQQNcZKZqJVXWvYxw3yFVec1QsfoX",
		"DzsRQ9TyubdKpYAEFNgTMPGQrzgZ1334HbnBiX1quKehA",
		"DormKcsSFKbmSa4es9arBXN3Mrdy5xPq1jtKtDwfoNEQr",
		"DitSkZRScZDPbFXxXJhhWXeHwjekGiXXeALME6tAAeVpV",
		"DW8TKCw8jsM7HNLMB7BRM69AtngYS3RzSgPwxp1bVAtTv",
		"DqPtdsgqWaKxgZidz35UPEA16Xrz3fBW1ECCBSESmckih",
		"DXqHn3HBcDPPKMaLHgxJx7s2gwhqR5ru3pzDzTps3kijx",
		"DtCgtuKo4t148tYeJ7zTj7FnniCEHafEGXtWKWCqTtJtx",
		"DpwmSgiSWWJW6QoNih5tRz4611Wa9sLziB8Xv37vrwRxM",
		"DjtZ3f8YQRpDz9iaM9zn3PW8Fm166E1xQ9ZsWPxwczBc9",
		"E36dedmpdFsbjG2HHEuiX9qSwve922CBJRuFrhSuPxaQH",
		"E44MZDxZHQQoBjBgT91FLWy4HLTtJrtbkrhHUMc7Qf1Nr",
		"E1DHTqKgoMfSzDgKMMikSLe1BL8Rpd3J9Gr76d2YraU5C",
		"E4eeX1uaB7tsqoyYJLxo3W8gMKpxfMTDYpTqLcxQoevw4",
		"DeAicjb4nfVeoqAoyZDiBKgKdBA22exsBMQKQDK6LVYjm",
		"DkhZjtQsgDoG39PTH75Mkmye3jG6ffYDPJebKuE3U4J9Y",
		"DmFPEEvBiym945Y8tYetLny9e1VPQveTWGKDCTmEU81ih",
		"DrXGFjJFfrYGAckgZLJ5PbtaLjtoQR5ysbYAXj99v4E5n",
		"DyRF6XcYJhZ4eVSkzpbdzUUaGBDA7VeiaYvysxKccdwi4",
		"DfJSgmQGSJDz9jeYGXKFRk4nnieN38Bj2oc6D6K35rFor",
		"Dg9CGUSDXEUxtYsxEwis4J1wMQdVsHduQ7KSHukq184ZG",
		"DVLiUbH3wm1NP3XgUGMpfx33gbD4yL7vA8i71eQg7fdhQ",
		"DdSqCoEQfpRsbR6DH7LoVfqu28mzdoxV72QLxC7fkKn4q",
		"Dkfr6GgdVp6CQnfrZm4QbkuCi4CH7PtAPoLL4ED8N3oWF",
		"DdeKrBxyMpxe3aLnfFQ2RicbLoL5PdBKeePY6pWfrLyfg",
		"DwHGLV9iMg2aZpBPyCXvUjZX6DUkoH8aH4P1Kz2J9XaoF",
		"DfUtzx5FUKS4cgWGzSP2fJ92VqU6dGbw4nzniW7GPNNQR",
		"Dt7DFQgRzg7uiAg7pgFaMk9AKAWNofqsypMeNPQX3pvNA",
		"DkcETc5nZSsU4vuDEfBao4YPjhM5eGkEZrFNH5Vpd7E6B",
		"DaYu1uQEKNxt8mYyz3riYZndAyZeVHFVtvpJTFbbXmzX7",
		"DoBmGKLLZsGqfZgJ5d8zmyZURLVCcDhH8Mmt7AtKV1QvR",
		"Df4XwZHTBZUsL9dVJgg8oTReHzaj9Rgb4aRKsDA5t5wJd",
		"DkTFSFuyyybcj2PVG8WsvFwcuaYzgLirb4wYKpG5BfFU9",
		"E2raiXkCGgV88NxAg4H7KJFKPzyxAc9cRDTb5tPpjjvo1",
		"DhVhbYztkzzdyvSFrPmVeBsCRsZ3WUYpMmQYg3o2YHTeZ",
		"E54opTdmgdhqBsySZc91rpCi914APvP6jQsCZxjN55ae7",
		"DnJN2QcrUnk4VHvhcyc1MxuXB4npeqQYJ2SKK5GwpHgtS",
		"DdR7vvR4WBxVC3FbzwjsoBYpbpa4aD2tMdbEgG4gsoCPZ",
		"DmWNMDTea54ibmCHWNr4Jepfv8pvkcpN2N8Cqj1mHpoxf",
	}

	ADDRESS          = addresses[rand.Intn(len(addresses))] // what a random address tho
	WORKERS          = 4                                               // concurrent workers to spawn
	SHARE_DIFFICULTY = 0                                               // 0 means that it'll be generated

	NODE_URL = "https://denaro-node.gaetano.eu.org/" // down 24/7
	POOL_URL = "https://denaro-pool.gaetano.eu.org/" // never down, nobody knows it
)

func getTransactionsMerkleTree(transactions []string) string {

	var fullData []byte

	for _, transaction := range transactions {
		data, _ := hex.DecodeString(transaction)
		fullData = append(fullData, data...)
	}

	hash := sha256.New()
	hash.Write(fullData)

	return hex.EncodeToString(hash.Sum(nil))
}

func checkBlockIsValid(blockContent []byte, shareChunk string, chunk string, idifficulty int, charset string, hasDecimal bool) (bool, bool) {

	hash := sha256.New()
	hash.Write(blockContent)

	blockHash := hex.EncodeToString(hash.Sum(nil))

	if hasDecimal {
		return strings.HasPrefix(blockHash, shareChunk), strings.HasPrefix(blockHash, chunk) && strings.Contains(charset, string(blockHash[idifficulty]))
	} else {
		return strings.HasPrefix(blockHash, shareChunk), strings.HasPrefix(blockHash, chunk)
	}
}

func worker(start int, step int, res MiningInfoResult, address string) {

	var difficulty float64 = res.Difficulty
	var idifficulty int = int(difficulty)
	var shareDifficulty int = idifficulty - 2

	_, decimal := math.Modf(difficulty)

	// if not 0 -> we've set our share_difficulty
	if SHARE_DIFFICULTY != 0 {
		shareDifficulty = SHARE_DIFFICULTY
	}

	lastBlock := res.LastBlock
	if lastBlock.Hash == "" {
		var num uint32 = 30_06_2005

		data := make([]byte, 32)
		binary.LittleEndian.PutUint32(data, num)

		lastBlock.Hash = hex.EncodeToString(data)
	}

	chunk := lastBlock.Hash[len(lastBlock.Hash)-idifficulty:]

	var shareChunk string

	if shareDifficulty > idifficulty {
		shareDifficulty = idifficulty
	}
	shareChunk = chunk[:shareDifficulty]

	charset := "0123456789abcdef"
	if decimal > 0 {
		count := math.Ceil(16 * (1 - decimal))
		charset = charset[:int(count)]
	}

	addressBytes := stringToBytes(address)
	t := float64(time.Now().UnixMicro()) / 1000000.0
	i := start
	a := time.Now().Unix()
	txs := res.PendingTransactionsHashes
	merkleTree := getTransactionsMerkleTree(txs)

	if start == 0 {
		log.Printf("Difficulty: %f\n", difficulty)
		log.Printf("Block number: %d\n", lastBlock.Id)
		log.Printf("Confirming %d transactions\n", len(txs))
	}

	var prefix []byte
	dataHash, _ := hex.DecodeString(lastBlock.Hash)
	prefix = append(prefix, dataHash...)
	prefix = append(prefix, addressBytes...)
	dataMerkleTree, _ := hex.DecodeString(merkleTree)
	prefix = append(prefix, dataMerkleTree...)
	dataA := make([]byte, 4)
	binary.LittleEndian.PutUint32(dataA, uint32(a))
	prefix = append(prefix, dataA...)
	dataDifficulty := make([]byte, 2)
	binary.LittleEndian.PutUint16(dataDifficulty, uint16(difficulty*10))
	prefix = append(prefix, dataDifficulty...)

	if len(addressBytes) == 33 {
		data1 := make([]byte, 2, 2)
		binary.LittleEndian.PutUint16(data1, uint16(2))

		oldPrefix := prefix
		prefix = data1[:1]
		prefix = append(prefix, oldPrefix...)
	}

	for {
		var _hex []byte

		found := true
		check := 5000000 * step

	checkLoop:
		for {
			if process, ok := processes.Load(start); !ok || !process.(Goroutine).Alive {
				return
			}

			_hex = _hex[:0]
			_hex = append(_hex, prefix...)
			dataI := make([]byte, 4)
			binary.LittleEndian.PutUint32(dataI, uint32(i))
			_hex = append(_hex, dataI...)

			shareValid, blockValid := checkBlockIsValid(_hex, shareChunk, chunk, idifficulty, charset, decimal > 0)

			if shareValid {
				var reqP Share

				req := POST(
					POOL_URL+"share",
					map[string]interface{}{
						"block_content":    hex.EncodeToString(_hex),
						"txs":              txs,
						"id":               lastBlock.Id + 1,
						"share_difficulty": difficulty,
					},
				)
				_ = json.Unmarshal(req.Body(), &reqP)

				if reqP.Ok {
					log.Println("SHARE ACCEPTED")
				} else {
					log.Println("SHARE NOT ACCEPTED")
					stopWorkers()
					return
				}
			}

			if blockValid {
				break checkLoop
			}

			i = i + step
			if (i-start)%check == 0 {
				elapsedTime := float64(time.Now().UnixMicro())/1000000.0 - t
				//log.Printf("Worker %d: %dk hash/s", start+1, i/step/int(elapsedTime)/1000)
        mu.Lock()
        f.WriteString(fmt.Sprintf("Worker %d: %dk hash/s\n", start+1, i/step/int(elapsedTime)/1000))
        mu.Unlock()

				if elapsedTime > 90 {
					found = false
					break checkLoop
				}
			}
		}

		if found {
			var reqP PushBlock

			log.Println(hex.EncodeToString(_hex))

			req := POST(
				POOL_URL+"push_block",
				map[string]interface{}{
					"block_content": hex.EncodeToString(_hex),
					"txs":           txs,
					"id":            lastBlock.Id + 1,
				},
			)
			_ = json.Unmarshal(req.Body(), &reqP)

			if reqP.Ok {
				log.Println("BLOCK MINED")
			}

			stopWorkers()
			return
		}
	}
}

func main() {

	flag.StringVar(&ADDRESS, "address", ADDRESS, "address that'll receive mining rewards")
	flag.IntVar(&WORKERS, "workers", WORKERS, "number of concurrent workers to spawn")
	flag.StringVar(&NODE_URL, "node", NODE_URL, "node to which we'll retrieve mining info")
	flag.StringVar(&POOL_URL, "pool", POOL_URL, "pool to which we'll mine on")
	flag.IntVar(&SHARE_DIFFICULTY, "share_difficulty", SHARE_DIFFICULTY, "pretty self descriptive")

	flag.Parse()

	var reqP MiningAddress

	req := GET(
		POOL_URL+"get_mining_address",
		map[string]interface{}{
			"address": ADDRESS,
		},
	)

	if err := json.Unmarshal(req.Body(), &reqP); err != nil {
		panic(err)
	}

	miningAddress := reqP.Address
	log.Println(miningAddress)

  f, _ = os.Create("out.txt")
  defer f.Close()

	for {
		log.Printf("Starting %d workers", WORKERS)

		var reqP MiningInfo

		req := GET(NODE_URL+"get_mining_info", map[string]interface{}{})
		_ = json.Unmarshal(req.Body(), &reqP)

		for _, i := range createRange(1, WORKERS) {
			log.Printf("Starting worker n.%d", i)
			go worker(i-1, WORKERS, reqP.Result, miningAddress)

			processes.Store(i-1, Goroutine{Id: i - 1, Alive: true, StartedAt: time.Now().Unix()})
		}

		elapsedSeconds := 0

	waitLoop:
		for allAliveWorkers() {
			time.Sleep(1 * time.Second)
			elapsedSeconds += 1

			if elapsedSeconds > 180 {
				stopWorkers()
				break waitLoop
			}
		}
	}
}

