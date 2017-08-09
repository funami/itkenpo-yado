import scrapy
import time

root = 'https://as.its-kenpo.or.jp'

class ItKenpoYadoSpider(scrapy.Spider):
    name = 'it_kenpo_yado_spider'
    start_urls = [root]
    custom_settings = {
        'USER_AGENT': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.75.14 (KHTML, like Gecko) Version/7.0.3 Safari/7046A194A',
    }
    def parse(self, response):
        next_page = root + response.css('.service_category a::attr(href)')[1].extract() 
        print(next_page)
        yield response.follow(next_page, callback=self.parse_yado)

    def parse_yado(self, response):
        for yado in response.css('li'):
            next_page = root + yado.css('a::attr(href)').extract_first()
            request = scrapy.Request(next_page, callback=self.parse_yado_monthly)
            request.meta['name'] = yado.css('a::text').extract_first() 
            yield request

    def parse_yado_monthly(self, response):
        for yado in response.css('li'):
            next_page = root + yado.css('a::attr(href)').extract_first()
            request = scrapy.Request(next_page, callback=self.parse_yado_detail)
            request.meta['name'] = response.meta['name'] 
            request.meta['name'] = yado.css('a::text').extract_first() 
            yield request

    def parse_yado_detail(self, response):
        dais = response.css('#apply_join_time option::attr(value)').extract()
        del dais[0]
        print(response.meta['name'], dais)
