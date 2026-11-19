
(ns ques2.core
  (:gen-class))

(defn -main [& args]

  ;; First Collection
  (def numbers [1 2 3 4 5 6])

  (println "First Collection:" numbers)

  ;; 1. map
  (println "Square:" (map #(* % %) numbers))

  ;; 2. filter
  (println "Even Numbers:" (filter even? numbers))

  ;; 3. reduce
  (println "Sum:" (reduce + numbers))

  ;; 4. take and drop
  (println "Take 3:" (take 3 numbers))
  (println "Drop 3:" (drop 3 numbers))

  ;; Threading Macro
  (println "Threading Result:"
           (->> numbers
                (filter even?)
                (map #(* % %))
                (take 2)))


  ;; Second Collection
  (def numbers2 [10 15 20 25 30])

  (println "\nSecond Collection:" numbers2)

  (println "Square:" (map #(* % %) numbers2))
  (println "Odd Numbers:" (filter odd? numbers2))
  (println "Sum:" (reduce + numbers2))
  (println "Take 2:" (take 2 numbers2))
  (println "Drop 2:" (drop 2 numbers2))

  ;; Threading Macro
  (println "Threading Result:"
           (->> numbers2
                (filter #(> % 15))
                (map inc))))
