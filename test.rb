cmd="curl http://localhost:8080/resize --output img.web"

threads = []

12.times do |n_th|
	threads << Thread.new do 
		1_000.times { |i| system(cmd) }
	end
end


threads.each { |th| th.join}

