# aguacate 🥑
A small app that parses a nmap XML output and generates a folder structure with some markdowns inside

## Build
Building is as easy as getting the repo and run `go build`. Alternatively you can download the latest build from the releases.

## Usage
Scan your assets with nmap and export to XML, something like:

```
nmap -sS -Pn -sV -v -O -T3 --open -iL ip.txt -oA scan
```

After, with the XML file execute _aguacate_:

```
./aguacate -nmap scan.xml
```