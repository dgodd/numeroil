require "roda"

class WordReduce
  def self.reduce(word)
    a = word.split('').map do |l|
      (l.upcase.ord - 'A'.ord) % 9 + 1
    end.to_a

    b = a.reduce(:+)
    c = b
    while c > 9 && ![11,22,33].include?(c) do
      c = c.to_s.split('').map do |d|
        d.to_i
      end.reduce(:+)
    end

    [b, c]
  end
end

class App < Roda
  top = "<html>
  <head>
  <link href='https://fonts.googleapis.com/css?family=Poiret+One' rel='stylesheet' type='text/css'><style>body{font-family: 'Poiret One', cursive;font-size:30px;background:linen;text-align:center;margin:100px 0;}</style></head><body>
          <form action='/word'>
          <label for='q'>Word</label>
          <input name='q'>
          <input type='submit'>
          </form><br><br>"
  bottom = "</body></html>"

  route do |r|
    r.root do
      top + bottom
    end

    r.get "word" do |word|
      a,b = WordReduce.reduce(r['q'])
      top + "Word: #{r['q']} <br> Big: #{a} <br> small: #{b}" + bottom
    end
  end
end

run App.freeze.app
