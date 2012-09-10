#/usr/bin/env ruby
# encoding: utf-8

require 'optparse'
require 'yaml'
require_relative 'app/gusgen_siren_bot'

AppName = 'gusgen_siren_bot'
Version = '0.1.0'

CONFIG_YAML = 'gusgen_siren_bot.yaml'

def print_help
  print <<HELP
#{AppName} #{Version}

Usage:
  #{File.basename($PROGRAM_NAME)} [option]

Options:
  -c FILE, --conf=FILE          Specify configuration file.
  -h, --help                    Show this help and exit.
  -v, --version                 Show version and exit.
HELP
end

OPTS = {
  conf: File.expand_path("../#{CONFIG_YAML}", __FILE__),
}

OptionParser.new { |opts|
  opts.on('-c', '--conf FILE', 'Specify configuration file.') do |v|
    OPTS[:conf] = v
  end

  opts.on_tail('-h', '--help', 'Show this help and exit.') do |v|
    print_help
    exit 0
  end
  opts.on_tail('-v', '--version', 'Show version and exit.') do |v|
    puts "#{AppName} #{Version}"
    exit 0
  end

  begin
    opts.parse!(ARGV)
  rescue => e
    puts e
    print_help
    exit 1
  end
}

if FileTest.exist?(OPTS[:conf])
  conf = YAML.load_file(OPTS[:conf])
else
  conf = {}
end

GusgenSirenBot.new(conf).run
