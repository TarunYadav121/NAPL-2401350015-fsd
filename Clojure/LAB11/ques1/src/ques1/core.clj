
(ns ques1.core
  (:gen-class))

(defn -main [& args]

  ;; 1. Java Static Methods
  (println "Absolute Value:" (Math/abs -25))
  (println "Current Time:" (System/currentTimeMillis))


  ;; 2. Create Java Objects
  (def arr (new java.util.ArrayList))
  (def hm (new java.util.HashMap))


  ;; 3. Invoke Instance Methods

  ;; ArrayList methods
  (.add arr "Apple")
  (.add arr "Banana")
  (.add arr "Mango")

  (println "\nJava ArrayList:" arr)
  (println "ArrayList Size:" (.size arr))
  (println "First Element:" (.get arr 0))

  ;; HashMap methods
  (.put hm "name" "Rahul")
  (.put hm "age" 20)

  (println "\nJava HashMap:" hm)
  (println "Name:" (.get hm "name"))
  (println "Map Size:" (.size hm))


  ;; 4. Convert Clojure Collection to Java Collection
  (def clj-vector [10 20 30 40])
  (def java-list (new java.util.ArrayList clj-vector))

  (println "\nClojure Vector:" clj-vector)
  (println "Java ArrayList:" java-list)


  ;; Convert Java Collection to Clojure Collection
  (def converted-vector (vec java-list))

  (println "Converted Clojure Vector:" converted-vector))
