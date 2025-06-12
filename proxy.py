import asyncio
from proxybroker import Broker

async def main():
    # Read proxies from file
    with open('proxy.txt', 'r') as f:
        lines = [line.strip() for line in f if line.strip()]
    proxies = []
    for line in lines:
        parts = line.split(':')
        if len(parts) == 2:
            host, port = parts
            proxies.append(f'{host}:{port}')
        elif len(parts) == 4:
            host, port, user, passwd = parts
            proxies.append(f'{user}:{passwd}@{host}:{port}')
        else:
            continue

    broker = Broker(None)

    # Start local proxy server on 127.0.0.1:8888
    await broker.serve(
        proxies=proxies,
        limit=100,
        addr='127.0.0.1',
        port=8888,
        max_conn=100
    )

if __name__ == '__main__':
    asyncio.run(main())
