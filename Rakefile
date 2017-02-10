

namespace :gcloud do

  task :activate do
    sh *%w( gcloud config configurations activate lazycache-appengine )
  end

  task :deploy => :activate do
    sh *%w( gcloud deploy )
  end

  task :open do
    sh *%w( open https://lazycache-dot-ferrous-ranger-158304.appspot.com )
  end

end

namespace :docker do

task :build do
  sh *%w( docker build -t go-lazycache-appengine:dev . )
end

task :run do
  sh *%w( docker run --rm --publish 8080:8080 go-lazycache-appengine:dev )
end

task :console do
  sh *%w( docker run --rm --tty --interactive --entrypoint /bin/bash --publish-all go-lazycache-appengine:dev )
end

end
