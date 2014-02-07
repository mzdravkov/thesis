func Deploy(name string, port uint16) (err error) {
  opts := strings.Join(dockerRunOptions(name, port), 
                       " ")
  args := strings.Join(dockerRunArguments(), " ")
  if err = makeTenantConfig(name); err != nil {
    return
  }
  if err = makePluginsDir(name); err != nil {
    return
  }
  dockerRunCmd := "docker run " + opts + " " + args
  cmd := exec.Command("sh", "-c", dockerRunCmd)
  if err = cmd.Run(); err != nil {
    return
  }
  locationOpts := locationOptions(port)
  if Config["nginx_use_locations"] == "true" {
    addLocation(name, locationOpts)
  } else {
    serverOpts := serverOptions(name)
    addServer(locationOpts, serverOpts)
  }
  reload := "sudo "+Config["nginx_bin"]+" -s reload"
  nginxReload := exec.Command("sh", "-c", reloadCmd)
  err = nginxReload.Run()
  return
}
