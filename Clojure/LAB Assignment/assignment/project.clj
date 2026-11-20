
(defproject assignment "0.1.0-SNAPSHOT"
  :description "Student Performance Data Pipeline"

  :url "https://example.com/assignment"

  :license {:name "EPL-2.0 OR GPL-2.0-or-later WITH Classpath-exception-2.0"
            :url "https://www.eclipse.org/legal/epl-2.0/"}

  :dependencies [[org.clojure/clojure "1.12.2"]
                 [ring/ring-core "1.12.2"]
                 [ring/ring-jetty-adapter "1.12.2"]
                 [ring/ring-json "0.5.1"]
                 [org.xerial/sqlite-jdbc "3.46.1.0"]]

  :main assignment.core
  :target-path "target/%s"

  :profiles {:uberjar {:aot :all
                       :jvm-opts ["-Dclojure.compiler.direct-linking=true"]}})
