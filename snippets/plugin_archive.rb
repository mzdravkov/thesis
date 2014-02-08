has_attached_file :archive, path: File.join(ENV['tenants_plugins_dir'],
                                            ':plugin_tenant_name',
                                            ':filename')

validates_attachment :archive, presence: true,
  content_type: {content_type: "application/x-gzip"},
  size: { in: 0..10.megabytes }

private

Paperclip.interpolates :plugin_tenant_name do |attachment|
  attachment.instance.tenant.name
end
