class AddTestTable < ActiveRecord::Migration[7.0]
  def change
    create_table :test do |t|
      t.string :name
      t.string :description
    end
  end
end
