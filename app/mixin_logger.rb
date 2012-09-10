# encoding: utf-8

require 'logger'

module MixinLogger
  attr_reader :logger

  def init_logger(conf)
    levels = { 'debug' => Logger::DEBUG, 'info'  => Logger::INFO, 'warn' => Logger::WARN,
               'error' => Logger::ERROR, 'fatal' => Logger::FATAL }

    if conf && conf['log'] && conf['log']['file']
      logdev = conf['log']['file']
      logdev = STDOUT if logdev == '-'

      @logger = Logger.new(logdev)
      if conf['log']['level'] && levels.key?(conf['log']['level'].downcase)
        @logger.level = levels[conf['log']['level'].downcase]
      end
    else
      init_dummy_logger
    end
  end

  private

  def init_dummy_logger
    @logger = Object.new
    class << @logger
      [:debug, :info, :warn, :error, :fatal].each do |method_name|
        class_eval(<<-METHOD_DEFN, __FILE__, __LINE__)
            def #{method_name}(msg=nil, &block)
              # do nothing
            end
        METHOD_DEFN
      end
    end
  end
end
