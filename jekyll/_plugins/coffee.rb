require 'coffee_script'
module Jekyll
  class CoffeeConverter < Converter
    safe true
    priority :low

    def matches(ext)
      ext =~ /coffee/i
    end

    def output_ext(ext)
      ".js"
    end

    def convert(content)
      CoffeeScript.compile(content)
    end
  end
end
