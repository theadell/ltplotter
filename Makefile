run-all:
	sh ./killports.sh
	cd go-services && make run-gateway &
	cd go-services && make run-exprplot &
	cd frontend && npm install && npm run dev &
	cd dataPlotService && dotnet run &
