# Roscomnadzor dump resolver

Library for parse and resolve all [Roskomnadzor registry](http://vigruzki.rkn.gov.ru/docs/description_for_operators_actual.pdf)

## Run CLI

```
go get github.com/velp/go-rknresolver
cd $GOPATH/src/github.com/velp/go-rknresolver
make build
./_output/rknresolver -d ./example/test.xml -o ./example/test.json
```

Example part of output JSON file:

```
$ head -n 25 ./example/test.json
{
  "content": [
    {
      "id": 347,
      "ip": [
        "109.69.58.58"
      ],
      "ipSubnet": null,
      "domain": [
        "video-one.com"
      ],
      "url": [
        "http://video-one.com/popular/Teen/1.html"
      ],
      "includeTime": "2012-11-18T15:17:51Z",
      "entryType": 1,
      "blockType": "",
      "urgencyType": 0,
      "hash": "",
      "decision": {
        "date": "2012-11-10T00:00:00Z",
        "number": "101-РИ",
        "org": "Роскомнадзор"
      }
    },
```

## CLI commands

```
$ rknresolver help
NAME:
   rknresolver - A new cli application

USAGE:
   rknresolver [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --dump value, -d value     path to XML dump file
   --out value, -o value      path to output JSON file
   --workers value, -w value  count workers in resolver (default: 100)
   --part value, -p value     part size (default: 1000)
   --dns value, -n value      DNS server for resolving (default: "127.0.1.1:53")
   --help, -h                 show help
   --version, -v              print the version
```

## Use lib

Please, see code *./cmd/rknresolver*.

## Performance

```
$ rknresolver -d dump.xml -o dump.json
2017/02/01 15:14:48 Create 100 workers
2017/02/01 15:14:48 Send all 59 jobs
2017/02/01 15:15:05 Progress (16.614021556s) 1000/58705
2017/02/01 15:15:08 Progress (20.343582879s) 2000/58705
2017/02/01 15:15:09 Progress (20.963941316s) 3000/58705
2017/02/01 15:15:10 Progress (21.462633989s) 3705/58705
2017/02/01 15:15:16 Progress (27.524748364s) 4705/58705
2017/02/01 15:15:16 Progress (28.206999422s) 5705/58705
2017/02/01 15:15:18 Progress (29.740935452s) 6705/58705
2017/02/01 15:15:24 Progress (35.891390401s) 7705/58705
2017/02/01 15:15:31 Progress (42.45396143s) 8705/58705
2017/02/01 15:15:34 Progress (45.699050356s) 9705/58705
2017/02/01 15:15:35 Progress (46.477247235s) 10705/58705
2017/02/01 15:15:36 Progress (48.353915628s) 11705/58705
2017/02/01 15:15:41 Progress (52.525873774s) 12705/58705
2017/02/01 15:15:45 Progress (56.915850542s) 13705/58705
2017/02/01 15:15:46 Progress (57.717593023s) 14705/58705
2017/02/01 15:15:48 Progress (1m0.138295469s) 15705/58705
2017/02/01 15:15:49 Progress (1m0.948825848s) 16705/58705
2017/02/01 15:15:51 Progress (1m2.65665571s) 17705/58705
2017/02/01 15:15:52 Progress (1m3.737170893s) 18705/58705
2017/02/01 15:15:53 Progress (1m4.526501382s) 19705/58705
2017/02/01 15:15:53 Progress (1m4.881689841s) 20705/58705
2017/02/01 15:15:54 Progress (1m5.589351155s) 21705/58705
2017/02/01 15:15:54 Progress (1m5.670823181s) 22705/58705
2017/02/01 15:15:54 Progress (1m6.345854681s) 23705/58705
2017/02/01 15:15:55 Progress (1m7.197175396s) 24705/58705
2017/02/01 15:15:55 Progress (1m7.197383475s) 25705/58705
2017/02/01 15:15:56 Progress (1m8.315804728s) 26705/58705
2017/02/01 15:15:59 Progress (1m10.504098514s) 27705/58705
2017/02/01 15:16:00 Progress (1m12.202005163s) 28705/58705
2017/02/01 15:16:00 Progress (1m12.273421722s) 29705/58705
2017/02/01 15:16:02 Progress (1m13.66466704s) 30705/58705
2017/02/01 15:16:02 Progress (1m14.326942084s) 31705/58705
2017/02/01 15:16:04 Progress (1m15.465708051s) 32705/58705
2017/02/01 15:16:04 Progress (1m16.013719797s) 33705/58705
2017/02/01 15:16:04 Progress (1m16.332604581s) 34705/58705
2017/02/01 15:16:06 Progress (1m18.23948293s) 35705/58705
2017/02/01 15:16:07 Progress (1m18.983214222s) 36705/58705
2017/02/01 15:16:07 Progress (1m19.273507152s) 37705/58705
2017/02/01 15:16:10 Progress (1m21.764725331s) 38705/58705
2017/02/01 15:16:10 Progress (1m21.885066584s) 39705/58705
2017/02/01 15:16:13 Progress (1m24.763693365s) 40705/58705
2017/02/01 15:16:13 Progress (1m24.88277923s) 41705/58705
2017/02/01 15:16:13 Progress (1m25.182346184s) 42705/58705
2017/02/01 15:16:14 Progress (1m26.100657681s) 43705/58705
2017/02/01 15:16:15 Progress (1m27.252123501s) 44705/58705
2017/02/01 15:16:15 Progress (1m27.409332166s) 45705/58705
2017/02/01 15:16:16 Progress (1m27.892980677s) 46705/58705
2017/02/01 15:16:17 Progress (1m29.0502714s) 47705/58705
2017/02/01 15:16:17 Progress (1m29.051682987s) 48705/58705
2017/02/01 15:16:17 Progress (1m29.293722251s) 49705/58705
2017/02/01 15:16:18 Progress (1m29.906967553s) 50705/58705
2017/02/01 15:16:19 Progress (1m31.262148151s) 51705/58705
2017/02/01 15:16:20 Progress (1m31.820795267s) 52705/58705
2017/02/01 15:16:21 Progress (1m32.888254996s) 53705/58705
2017/02/01 15:16:24 Progress (1m36.080071665s) 54705/58705
2017/02/01 15:16:25 Progress (1m36.801283801s) 55705/58705
2017/02/01 15:16:27 Progress (1m39.401506215s) 56705/58705
2017/02/01 15:16:28 Progress (1m39.74895187s) 57705/58705
2017/02/01 15:16:30 Progress (1m41.580335237s) 58705/58705
```