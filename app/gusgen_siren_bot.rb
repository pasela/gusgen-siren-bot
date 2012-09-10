# encoding: utf-8

require 'vanadiel/time'
require_relative 'mixin_logger'

class GusgenSirenBot
  include MixinLogger

  DATE_FORMAT = '%Y-%m-%d'

  def initialize(conf)
    @conf = conf
    init_logger(@conf)
  end

  def run
    vana_time = Vanadiel::Time.now
    logger.info "startup at #{vana_time}(#{vana_time.to_earth_time})"

    update_last_date(vana_time)
    loop do
      vana_time = Vanadiel::Time.now
      logger.debug "current time = #{vana_time}(#{vana_time.to_earth_time})"

      if date_changed?(vana_time)
        update_last_date(vana_time)
        onchangedate(vana_time)
      end

      tomorrow = vana_time + (24 * 60 * 60)
      next_time = Vanadiel::Time.new(tomorrow.year, tomorrow.mon, tomorrow.day)
      logger.info "next time = #{next_time}(#{next_time.to_earth_time})"
      sleep(next_time.to_earth_time - vana_time.to_earth_time)
    end
  end

  private

  def onchangedate(vana_time)
    puts "siren at #{vana_time}(#{vana_time.to_earth_time})"
  end

  def update_last_date(vana_time)
    @last_date = vana_time.strftime(DATE_FORMAT)
  end

  def date_changed?(vana_time)
    vana_date = vana_time.strftime(DATE_FORMAT)
    (@last_date.nil? || vana_date > @last_date) && (vana_time.hour == 0 && vana_time.min == 0)
  end
end
