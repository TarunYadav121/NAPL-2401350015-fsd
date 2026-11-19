
(ns immutable-ds-project.core)

(defn -main [& args]

  ;; 1. LIST
  (def my-list '(10 20 30))

  (println "Original List:" my-list)
  (println "First:" (first my-list))
  (println "Rest:" (rest my-list))

  (def new-list (conj my-list 5))
  (println "New List:" new-list)
  (println "Original List:" my-list)


  ;; 2. VECTOR
  (def my-vector [10 20 30])

  (println "\nOriginal Vector:" my-vector)
  (println "Get index 1:" (get my-vector 1))
  (println "Count:" (count my-vector))

  (def new-vector (conj my-vector 40))
  (println "New Vector:" new-vector)
  (println "Original Vector:" my-vector)


  ;; 3. MAP
  (def my-map {:name "Rahul" :age 20})

  (println "\nOriginal Map:" my-map)
  (println "Get name:" (get my-map :name))

  (def new-map (assoc my-map :city "Delhi"))
  (println "New Map:" new-map)

  (def updated-map (dissoc my-map :age))
  (println "After dissoc:" updated-map)
  (println "Original Map:" my-map)


  ;; 4. SET
  (def my-set #{10 20 30})

  (println "\nOriginal Set:" my-set)
  (println "Contains 20:" (contains? my-set 20))

  (def new-set (conj my-set 40))
  (println "New Set:" new-set)

  (def updated-set (disj my-set 20))
  (println "After disj:" updated-set)
  (println "Original Set:" my-set))
