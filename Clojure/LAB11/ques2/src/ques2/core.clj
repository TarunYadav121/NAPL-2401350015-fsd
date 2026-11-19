
(ns ques2.core
  (:gen-class))

;; Custom exception function
(defn check-age [age]
  (if (< age 18)
    (throw (ex-info "Age must be 18 or above"
                    {:age age :type "Invalid Age"}))
    "Valid Age"))

(defn -main [& args]

  ;; 1. Arithmetic Exception
  (println "\n1. Arithmetic Exception:")

  (try
    (println (/ 10 0))
    (catch ArithmeticException e
      (println "Error:" (.getMessage e)))
    (finally
      (println "Arithmetic Finally Block")))


  ;; 2. Null Pointer Exception
  (println "\n2. Null Pointer Exception:")

  (try
    (.length ^String (identity nil))
    (catch NullPointerException e
      (println "Error: Null value found"))
    (finally
      (println "Null Finally Block")))


  ;; 3. Custom Exception
  (println "\n3. Custom Exception:")

  (try
    (println (check-age 15))
    (catch clojure.lang.ExceptionInfo e
      (println "Error:" (ex-message e))
      (println "Exception Data:" (ex-data e))
      (println "Recovery: Please enter a valid age"))
    (finally
      (println "Custom Finally Block")))


  (println "\nProgram continues normally!"))
