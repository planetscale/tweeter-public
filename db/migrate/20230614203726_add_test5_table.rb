class AddTest5Table < ActiveRecord::Migration[7.0]
  def change
    create_table :test5 do |t|
      t.string :name
    end
  end
end
