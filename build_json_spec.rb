require 'nokogiri'
require 'json'


doc = Nokogiri::HTML(open('spec.html'))

spec = doc.css('.example').map do |example|
  {
    Name: example.css('div.examplenum').text.strip,
    Markdown: example.css('pre.markdown').text.gsub("→", "\t"),
    HTML: example.css('pre.html').text
  }
end

puts JSON.pretty_generate({Examples: spec})
File.open("spec.json", 'w') { |file| file.write(JSON.pretty_generate({Examples: spec})) }
