# How to Run
go run main.go -wordlist='pass.txt' -hash=6a85dfd77d9cb35770c9dc6728d73d3f
# For verbos output
go run main.go -wordlist='pass.txt' -hash=6a85dfd77d9cb35770c9dc6728d73d3f -out=verbose.txt

# for change algorithm
go run main.go -wordlist='pass.txt' -hash=aa1c7d931cf140bb35a5a16adeb83a551649c3b9 -out='verbose.txt' -type=sha1

go run main.go -wordlist='pass.txt' -hash=485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38 -out='verbose.txt' -type=sha512

go run main.go -wordlist='pass.txt' -hash=6a85dfd77d9cb35770c9dc6728d73d3f -out=verbose.txt -type=md5

