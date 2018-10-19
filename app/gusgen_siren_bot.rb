# encoding: utf-8

require 'vanadiel/time'
require 'twitter'
require_relative 'mixin_logger'

class GusgenSirenBot
  include MixinLogger

  DATE_FORMAT = '%Y-%m-%d'

  def initialize(conf)
    @conf = conf
    init_logger(@conf)
    begin
      init_twitter(@conf)
    rescue => e
      logger.error "failed to initialize twitter: #{e}"
      raise
    end
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

        limit_time = Vanadiel::Time.new(vana_time.year, vana_time.month, vana_time.mday, 1)
        begin
          try_until(vana_time, limit_time) do |vt|
            onchangedate(vt)
          end
        rescue => e
          logger.error "retry over: #{e}"
        end
      end

      tomorrow = vana_time + (24 * 60 * 60)
      next_time = Vanadiel::Time.new(tomorrow.year, tomorrow.mon, tomorrow.day)
      interval = next_time.to_earth_time - vana_time.to_earth_time
      if interval > (30 * 60)
        logger.info "next time = #{next_time}(#{next_time.to_earth_time})"
      end
      sleep(interval >= 1 ? interval / 2 : interval)
    end
  end

  private

  def init_twitter(conf)
    @twitter = Twitter::REST::Client.new do |config|
      config.consumer_key       = conf['oauth']['consumer_key']
      config.consumer_secret    = conf['oauth']['consumer_secret']
      config.access_token        = conf['oauth']['access_token']
      config.access_token_secret = conf['oauth']['access_token_secret']
    end

    user = @twitter.verify_credentials
    logger.info "verify_credentials ...ok (id=#{user.id}, screen_name=#{user.screen_name})"
  end

  def onchangedate(vana_time)
    begin
      tweet = @twitter.update(create_tweet(vana_time))
      logger.debug "tweet: id=#{tweet.id}, created_at=#{tweet.created_at}, text=\"#{tweet.text}\""
    rescue => e
      logger.error "failed to tweet #{vana_time}/#{vana_time.to_earth_time}: #{e}"
      raise
    end
  end

  def create_tweet(vana_time)
    vana_info = "(%s)" % format_vana_time(vana_time)
    "ウゥーーーーーーーーーーーーーーー\n#{vana_info}"
  end

  def format_vana_time(vana_time)
    str = '天晶暦%s %s %s' % [vana_time.strftime('%Y/%m/%d %H:%M:%S'),
                              Vanadiel::Day::DAYNAMES_JA[vana_time.wday],
                              Vanadiel::Moon::MOONNAMES_JA[vana_time.moon_age]]
    str
  end

  def update_last_date(vana_time)
    @last_date = vana_time.strftime(DATE_FORMAT)
  end

  def date_changed?(vana_time)
    vana_date = vana_time.strftime(DATE_FORMAT)
    (@last_date.nil? || vana_date > @last_date) && (vana_time.hour == 0 && vana_time.min == 0)
  end

  def try_until(vt, limit, sleep_time = 60)
    begin
      yield vt
    rescue
      if vt < limit
        sleep sleep_time
        vt = Vanadiel::Time.now
        retry if vt <= limit
      else
        raise
      end
    end
  end
end
