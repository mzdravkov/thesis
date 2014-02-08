class Log < Plugman::PluginBase
  def on_deploy tenant
    `cat '#{tenant.name}: #{tenant.user.email}' >> log`
  end
end
