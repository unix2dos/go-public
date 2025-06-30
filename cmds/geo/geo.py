import maxminddb
import requests
from collections import defaultdict

# 下载 mmdb 文件
url = "https://github.com/P3TERX/GeoLite.mmdb/raw/download/GeoLite2-Country.mmdb"
response = requests.get(url)
with open("GeoLite2-Country.mmdb", "wb") as f:
    f.write(response.content)

# 读取并提取所有国家
reader = maxminddb.open_database('GeoLite2-Country.mmdb')
countries = set()

# 遍历数据库中的所有记录
for record in reader:
    if record and 'country' in record:
        country_info = record['country']
        if 'names' in country_info:
            # 获取英文名称
            if 'en' in country_info['names']:
                countries.add(country_info['names']['en'])
            # 也可以获取中文名称
            if 'zh-CN' in country_info['names']:
                countries.add(country_info['names']['zh-CN'])

# 打印所有国家（按字母排序）
for country in sorted(countries):
    print(country)

reader.close()
